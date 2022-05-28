window.addEventListener('DOMContentLoaded', function(){

	// パスワード入力欄とボタンのHTMLを取得
	let btn_passview = document.getElementById('btn_passview');
	let Password = document.getElementById('superUserPassword');

	// ボタンのイベントリスナーを取得
	btn_passview.addEventListener("click", (e) => {

		// ボタンの通常の動作をキャンセル
		e.preventDefault();

		// パスワード入力欄のtype属性を確認
		if ( Password.type === 'password' ){

			// パスワードを表示する
			Password.type = 'text';
			btn_passview.textContent = '非表示';
		} else {

			// パスワードを非表示にする
			Password.type = 'password';
			btn_passview.textContent = '表示';
		}
	});
});

new Vue({
	// 「el」プロパティーで、Vueの表示を反映する場所を定義
	el: '#app',

	// data オブジェクトの定義
	data: {
		// Super User ID
		superUserID: '',
		// Super User Password
		superUserPassword: '',
		// Display Message
		displayMessage: '',
		// 画面表示フラグ true:表示／false非表示
		isDisplay: false,
	},

	// インスタンス作成時の処理
	created: function(){
		this.openSuperUserPassword()
	},

	// メソッド定義
	methods: {
		doUpdateSuperUser() {
			// サーバへ送信するパラメータ
			const params = new URLSearchParams();
			params.append('superUserID', this.superUserID);
			params.append('superUserPassword', this.superUserPassword);

			axios.post('/InsertSuperUserPassword', params)
			.then(response => {
				if (response.status != 200) {
					this.displayMessage = '管理者情報の更新に失敗しました'
				} else {
					this.displayMessage = '管理者情報の更新に成功しました'
				}
			})
		},
		// 管理者専用パスワード処理
		openSuperUserPassword() {
			inputPassword = prompt('管理者専用パスワードを入力してください');

			// キャンセルボタン押下時はエラーページへ
			if (inputPassword == null){
				let url = './superusererror.html';
				location.href = url;
			}

			axios.get('/superUserPasswordCollation', {
 				params: {
					productPassword: inputPassword
				}
			})
			.then(response => {
				if ( response.status == 200 ){
					// パスワードが一致している場合のみ、当ページの内容表示
					this.isDisplay = true
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
					throw new Error('openSuperUserPassword Response Error')
				}
			})
		}
	}
})

