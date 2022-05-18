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

	},

	// インスタンス作成時の処理
	created: function() {
		this.MessageBoardOpen()
	},

	// メソッド定義
	methods: {
		MessageBoardOpen() {
			let InputPassword = prompt('管理者専用パスワードを入力してください');

			// キャンセルボタン押下時はエラーページへ
			if (InputPassword == null){
				this.isMessageBoard = false;
				let url = './superusererror.html';
				location.href = url;
			}

			// サーバーにパスワードが一致しているか問い合わせる
			axios.get('/superUserPasswordCollation', {
 				params: {
					productPassword: InputPassword
				}
			})
			.then(response => {
				if ( response.status === 200 ){
					// パスワードが一致している場合はページを表示
					this.isMessageBoard = true;
				}
			})
			.catch( function(error) { 
				if ( error.response.status === 401 ){
					// パスワード不一致時はエラーページへ
					this.isMessageBoard = false;
					let url = './superusererror.html';
					location.href = url;
				}
			})
		}
	}
})
