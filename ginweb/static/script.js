// 获取Stories和Duration数据
// {{ . | jsonify }} 是服务端模板的占位符，会被实际的数据替换
// const { Stories: stories, Duration: duration } = {{ . | jsonify }};
let stories = []; // 存储故事的数组

// 创建WebSocket连接
const socket = new WebSocket("ws://localhost:8079/ws");
// 当连接建立时触发
socket.onopen = function(event) {
  console.log("WebSocket连接已建立");
};
// 当发生错误时触发
socket.onerror = function(event) {
  console.error("WebSocket错误:", event);
};

// 当接收到WebSocket消息时的处理函数
socket.onmessage = function (event) {
  const message = event.data;
  console.log("接收到消息:", message);
  // 在这里处理接收到的消息
  const receivedStories = JSON.parse(event.data); // 解析收到的故事数据
  stories = receivedStories; // 更新故事数组
  console.log("Stories:", stories);
  updateStories(); // 更新故事显示
};

// 当WebSocket连接关闭时的处理函数
socket.onclose = function (event) {
  console.log("WebSocket connection closed: ", event);
};

// 更新故事显示的函数
function updateStories() {
  const storiesContainer = document.getElementById("stories-container");
  storiesContainer.innerHTML = ""; // 清空故事容器

  for (let i = 0; i < stories.length; i++) {
    const story = stories[i];
    const storyDiv = document.createElement("div");
    storyDiv.className = "story";
    storyDiv.innerHTML = `
      <h2>${story.Title}</h2>
      <p>${story.Content}</p>
    `;
    storiesContainer.appendChild(storyDiv); // 添加故事到容器中
  }

  // 开始轮播故事
  carouselStories(5000);// 设置轮播间隔时间为5000秒，作废，设久一点
}

// 轮播故事函数
function carouselStories(duration) {
  console.log("开始轮播故事");
  const carouselElements = document.querySelectorAll('.carousel .story');
  let currentStoryIndex = 0;

  console.log("故事数量:", carouselElements.length);
  function showStory(index) {
    carouselElements.forEach((element, i) => {
      if (i === index) {
        element.style.display = 'block';
      } else {
        element.style.display = 'none';
      }
    });
  }

  function nextStory() {
    currentStoryIndex = (currentStoryIndex + 1) % stories.length;
    showStory(currentStoryIndex);
  }

  // 显示第一个故事
  showStory(currentStoryIndex);

  // 开始轮播故事
  setInterval(nextStory, duration * 1000);
}


