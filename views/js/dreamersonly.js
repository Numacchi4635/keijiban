window.addEventListener('DOMContentLoaded', function(){

	// パスワード入力欄とボタンのHTMLを取得
	let btn_passview = document.getElementById('btn_passview');
	let Password = document.getElementById('Password');

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
		UserID: '',
		// Super User Password
		Password: '',
		// Display Message
		DisplayMessage: ''
	},

	// メソッド定義
	methods: {
		doUpdateSuperUser() {
console.log('doUpdataSuperUser Start');
			// サーバへ送信するパラメータ
			const params = new URLSearchParams();
			params.append('UserID', this.UserID);
			params.append('Password', this.Password);

			axios.post('/InsertSuperUserPassword', params)
			.then(response => {
				if (response.status != 200) {
					throw new Error('UpsateSuperUser Response Error')
					this.DisplayMessage = '管理者情報の更新に失敗しました'
				} else {
					this.DisplayMessage = '管理者情報の更新に成功しました'
				}
			})
		}
	}
})

