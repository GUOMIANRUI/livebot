/*
添加一个全局变量来保存故事。
创建一个WebSocket处理程序来建立WebSocket连接。
当建立新的WebSocket连接时，向客户端发送故事。
当接收到新的故事时，更新全局变量并将更新后的故事广播给所有连接的客户端。
*/
package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	pb "webtest/stub/gochat/proto"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
)

type Story struct {
	Title   string // 故事标题
	Content string // 故事内容
}

var (
	stories          []Story
	upgrader         = websocket.Upgrader{}
	connectedSockets []*websocket.Conn
)

func main() {
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/static", "./static")

	// 设置模板文件目录
	r.LoadHTMLGlob("templates/*.html")

	templatePath := "templates/index.html" // 相对于Go文件的路径
	_, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	// 轮播页面路由
	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{})

		stories, err = readStoriesFromFile("gushi.txt")
		if err != nil {
			c.AbortWithError(404, err)
			return
		}

		fmt.Printf("stories: %v\n", stories)
	})

	// trpc client
	proxy := pb.NewGoChatServiceClientProxy(
		client.WithTarget("ip://127.0.0.1:8000"),
		client.WithProtocol("trpc"),
	)
	ctx := trpc.BackgroundContext()
	var voicefilename string

	// WebSocket路由
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Failed to upgrade to WebSocket:", err)
			return
		}

		connectedSockets = append(connectedSockets, conn)

		log.Printf("Connected to %s\n", conn.RemoteAddr())
		log.Printf("Sending %d stories to %s\n", len(stories), conn.RemoteAddr())
		// err = conn.WriteJSON(stories)
		// if err != nil {
		// 	log.Println("Failed to send stories:", err)
		// }
		storyarr := make([]Story, 1)
		title := make(map[string]bool)
		timeuse := int32(0)

		// 将每个故事单独发送给客户端
		for p0, story := range stories {
			timeuse = 0
			pp := 0
			// 先读标题，并计算时间
			if _, ok := title[story.Title]; !ok {
				title[story.Title] = true
				req := &pb.GoChataudioRequest{}
				voicefilename = strconv.Itoa(p0) + "-" + "title" + "-web.mp3"
				req.Filename = voicefilename
				req.Msg = story.Title
			retryaudiot:
				reply, err := proxy.GoChatAudio(ctx, req)
				if err != nil {
					if pp <= 4 {
						pp++
						goto retryaudiot
					}
					log.Fatalf("err: %v", err)
				}
				timeuse += reply.Filelongth + 2
				// 等待时间
				time.Sleep(time.Duration(timeuse) * time.Second)
				timeuse = 0
			}
			// 如果story内容按空格拆成多个片段，如果片段长于4，每4段发送一次，并返回时间
			storyspl := strings.Split(story.Content, " ")
			if len(storyspl) > 4 {
				for i := 0; i < len(storyspl); i += 4 {
					if i+4 > len(storyspl) {
						// 将剩余的片段组合成一个字符串
						content := strings.Join(storyspl[i:], " ")
						storyarr[0] = Story{Title: story.Title, Content: content}
					} else {
						// 将当前4个片段组合成一个字符串
						content := strings.Join(storyspl[i:i+4], " ")
						storyarr[0] = Story{Title: story.Title, Content: content}
					}

					// 发送故事给客户端
					err := conn.WriteJSON(storyarr)
					if err != nil {
						log.Println("Failed to send story:", err)
						break
					}

					// 请求server，转语音
					voicefilename = strconv.Itoa(p0) + "-" + strconv.Itoa(i) + "-" + "conent" + "-web.mp3"
					req := &pb.GoChataudioRequest{}
					req.Filename = voicefilename
					req.Msg = storyarr[0].Content
					pp = 0
				retryaudiot2:
					reply, err := proxy.GoChatAudio(ctx, req)
					if err != nil {
						if pp <= 4 {
							pp++
							goto retryaudiot2
						}
						log.Fatalf("err: %v", err)
					}
					timeuse += reply.Filelongth + 2

					// 等待时间
					time.Sleep(time.Duration(timeuse) * time.Second)
					timeuse = 0
				}
			} else {
				storyarr[0] = story
				err := conn.WriteJSON(storyarr)
				if err != nil {
					log.Println("Failed to send story:", err)
					break
				}
				// 请求server，转语音
				voicefilename = strconv.Itoa(p0) + "-short-" + "conent" + "-web.mp3"
				req := &pb.GoChataudioRequest{}
				req.Filename = voicefilename
				req.Msg = storyarr[0].Content
				pp = 0
			retryaudiot3:
				reply, err := proxy.GoChatAudio(ctx, req)
				if err != nil {
					if pp <= 4 {
						pp++
						goto retryaudiot3
					}
					log.Fatalf("err: %v", err)
				}
				timeuse += reply.Filelongth + 2
				// 等待时间
				time.Sleep(time.Duration(timeuse) * time.Second)
				timeuse = 0
			}
		}

		// 服务器将等待客户端发送新的故事，并将每个新的故事追加到全局故事变量中
		for {
			var newStories []Story
			err := conn.ReadJSON(&newStories)
			if err != nil {
				log.Println("Failed to read stories from client:", err)
				break
			}

			fmt.Printf("newStories: %v\n", newStories)
			// 更新全局故事变量
			stories = newStories

			// 将更新后的故事广播给所有连接的客户端
			broadcastStories()
		}

		// 清理关闭的WebSocket连接
		for i, socket := range connectedSockets {
			if socket == conn {
				connectedSockets = append(connectedSockets[:i], connectedSockets[i+1:]...)
				break
			}
		}
	})

	r.Run(":8079")
}

func broadcastStories() {
	for _, conn := range connectedSockets {
		err := conn.WriteJSON(stories)
		if err != nil {
			log.Println("Failed to send stories to client:", err)
		}
	}
}

func readStoriesFromFile(filename string) ([]Story, error) {
	var stories []Story

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentStory Story
	isContentLine := false

	for scanner.Scan() {
		line := scanner.Text()
		// log.Printf("line=%s\n", line)

		if strings.HasPrefix(line, " ") {
			// 如果行以空格开头，则为内容行
			if isContentLine {
				// 如果是内容行，则将内容添加到当前故事的 Content 字段中
				currentStory.Content += line + "\n"
			}
		} else {
			// 如果行不以空格开头，则为标题行
			if currentStory.Title != "" {
				// 如果当前故事的标题不为空，则将当前故事添加到故事列表中
				stories = append(stories, currentStory)
			}

			// 创建新的故事，将标题行作为新故事的标题
			currentStory = Story{
				Title:   line,
				Content: "",
			}
			isContentLine = true
		}
	}

	if currentStory.Title != "" {
		// 将最后一个故事添加到故事列表中
		stories = append(stories, currentStory)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stories, nil
}
