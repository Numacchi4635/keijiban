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
		PublicMode: '',
		// タイトル
		title: '',
		// メッセージボードリンクボタン true:表示／false非表示
		isButton: false,
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
			axios.get('/responseServerEnv')
			.then(response => {
				if (response.status != 200) {
					throw new Error('responseServerEnv Response Error')
				} else {
					var resultResponse = response.data

					// 取得した環境変数毎にタイトル変更
					if (resultResponse.PublicMode === 'public'){
						this.title = '管理者専用メッセージ投稿ページ'
					} else if (resultResponse.PublicMode === 'private'){
						this.title = '🐹🍎ゆゆこ🍎🐹専用メッセージ投稿ページ'
					} else {
						this.title = '管理者専用メッセージ投稿ページ'
					}

					// 環境変数privateの時のみ、メッセージボードの内容へのリンクを表示
					if (resultResponse.PublicMode == 'private'){

						// メッセージリンクボタン表示をtrueにする
						this.isButton = true;

						// メッセージボードへのリンクボタンタグ設定
//						let displayMessageBoardTag = '<p><button type="button" onclick="window.open(\'./2021autmnmessage.html\')">2021年秋のメッセージボードへ</button></p><p><button type="button" onclick="window.open(\'./2022wintermessage.html\')">2022年冬のメッセージボードへ</button></p>'
						// メッセージボード表示
//						let messageBoard = document.getElementById('messageboardinfo');
//						messageBoard.insertAdjacentHTML('afterend', displayMessageBoardTag);
					}

				}
			})
		}
	}
})
