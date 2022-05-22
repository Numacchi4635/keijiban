new Vue({
	// 「el」プロパティーで、Vueの表示を反映する場所を定義
	el: '#app',

	// data オブジェクトの定義
	data: {
		// 宛先
		productName: '',
		// メッセージ
		productMessage: '',
		// 管理者パスワード
		superUserPassword: '',
		// publicMode(サーバー側の環境変数)
		PublicMode: '',
		// タイトル
		title: '',
		// メッセージボードリンクボタン true:表示／false非表示
		isButton: false,
		// 画面表示フラグ true:表示／false非表示
		isDisplay: false,
	}, 

	// インスタンス作成時の処理
	created: function() {
		this.openSuperUserPassword()
	},

	// メソッド定義
	methods: {

		doAddProduct() {
			// サーバへ送信するパラメータ
			const params = new URLSearchParams()
			params.append('productName', this.productName)
			params.append('productMessage', this.productMessage)
			params.append('superUserPassword', this.superUserPassword)

			axios.post('/addProduct', params)
			.then(response => {
				if (response.status != 200) {
					throw new Error('addProduct Response Error')
				} else {
					// index.htmlに戻る
					location.href = './index.html';
				}
			})
			.catch(function(error){
				if (error.response.status == 401){
					// パスワードが一致していない場合はエラーページへ
					this.isKeijibanDisplay = false
					let url = './superusererror.html';
					location.href = url;
				} else {
					throw new Error('addProduct Response Error')
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
					}
				}
			})
		},
		// 管理者専用パスワード処理
		openSuperUserPassword() {
			this.superUserPassword = prompt('管理者専用パスワードを入力してください');

			// キャンセルボタン押下時はエラーページへ
			if (this.superUserPassword == null){
				let url = './superusererror.html';
				location.href = url;
			}

			axios.get('/superUserPasswordCollation', {
 				params: {
					productPassword: this.superUserPassword
				}
			})
			.then(response => {
				if ( response.status == 200 ){
					// パスワードが一致している場合のみ、当ページの内容表示
					this.isDisplay = true
					this.responseServerEnv()
				}
			})
			.catch(function(error){
				if ( error.response.status == 401) {
					// パスワードが一致していない場合はエラーページへ
					this.isDisplay = false
					let url = './superusererror.html';
					location.href = url;
				} else {
					// 上記以外のエラーの場合
					throw new Error('fetchProduct Response Error')
				}
			})
		}
	}
})
