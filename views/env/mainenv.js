require('dotenv').config()
const env = process.env
console.log(env.PUBLIC_MODE)

if (env.PUBLIC_MODE == "public") {
	document.querySelector("h1").innerText = "パスワード認証付き掲示板";
} else {
	document.querySelector("h1").innerText = "🐹🍎ゆゆこ🍎🐹ファミリーボード返信掲示板";
}
