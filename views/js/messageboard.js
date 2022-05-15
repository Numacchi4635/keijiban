new Vue({
	// 「el」プロパティーで、Vueの表示を反映する場所を定義
	el: '#app',

	// data オブジェクトの定義
	data: {
		// スーパーユーザーパスワード
		Password: '',

		// パスワードが不一致時のエラー表示
		ErrorMessage: '',

		// メッセージボード true:表示／false:非表示
		isMessageBoard: false,

		// エラーメッセージ true:表示／false:非表示
		isErrorMessage: true,
	},

	// インスタンス作成時の処理
	created: function() {
		this.MessageBoardOpen()
	},

	// メソッド定義
	methods: {
		MessageBoardOpen() {
console.log(this.isMessageBoard)
			let InputPassword = prompt('管理者専用パスワードを入力してください');

			// サーバーにパスワードが一致しているか問い合わせる
			const params = new URLSearchParams()
			params.append('superUserPassword', InputPassword)
			axios.post('/superUserPasswordCollation', params)
			.then(response => {
				if ( response.status === 200 ){
					// パスワードが一致している場合はページを表示
console.log('パスワード一致')
					this.isMessageBoard = true;
					this.isErrorMessage = false;
				} else {
console.log('パスワード不一致')
					// パスワードが一致していない場合はエラー表示する
					this.ErrorMessage = 'パスワードが一致していないので当ページは表示できません'
					this.isMessageBoard = false;
					this.isErrorMessage = true;
console.log(this.ErrorMessage)
				}
			})
		}
	}
})
