new Vue({
	// 「el」プロパティーで、Vueの表示を反映する場所を定義
	el: '#app',

	// data オブジェクトの定義
	data: {
		// 宛先
		productName: '',
		// メッセージ
		productMessage: '',
		// PUBLIC_MODE(サーバー側の環境変数)
		PublicMode: ''
	}, 

	// インスタンス作成時の処理
	created: function() {
		this.responseServerEnv()
	},

	// メソッド定義
	methods: {

		doAddProduct() {
			// サーバへ送信するパラメータ
			const params = new URLSearchParams()
			params.append('productName', this.productName)
			params.append('productMessage', this.productMessage)

			axios.post('/addProduct', params)
			.then(response => {
				if (response.status != 200) {
					throw new Error('addProduct Response Error')
				} else {
					// index.htmlに戻る
					location.href = './index.html';
				}
			})
		},
		responseServerEnv() {
			window.onload = function(){
				axios.get('/responseServerEnv')
				.then(response => {
					if (response.status != 200) {
						throw new Error('responseServerEnv Response Error')
					} else {
						var resultResponse = response.data

						// 取得した環境変数毎にタイトル変更
						let displayTitleTag;
						if (resultResponse.PublicMode == 'public'){
							displayTitleTag = '<h1>管理者専用メッセージ投稿ページ</h1>'
						} else {
							displayTitleTag = '<h1>🐹🍎ゆゆこ🍎🐹専用メッセージ投稿ページ</h1>'
						}
						// タイトル表示
						let element = document.getElementById('titleinfo');
						element.insertAdjacentHTML('afterend',  displayTitleTag);
					}
				})
			}
		}
	}
})
