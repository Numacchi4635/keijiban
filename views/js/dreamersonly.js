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

