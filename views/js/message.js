// Vue
var app = new Vue({
	el: "#app",
	data: {
		productID: '',
		productName: '',
		productMessage: '',
		ProductPassword: '',
	},
	// インスタンス作成時の処理
	created: function(){
		this.GetProductData()
	},

	methods: {
		GetProductData(){
			// パスワード入力
			let password = prompt('パスワードを入力してください');


			// URLパラメータ取得
			let url = new URL(window.location.href);
			let params = url.searchParams;
			let Id = params.get('id');

			axios.get('/fetchProduct', {
				params: {
					productID: Id,
					productPassword: password
				}
			})
			.then(response => {
				if (response.status == 200){
					var resultProducts = response.data

					// 掲示板表示タグ生成
					this.productName = resultProducts.Name;
					this.productMessage = resultProducts.Message;
				}
			})
			.catch(function(error){
				if (error.response.status == 401) {
					// パスワードが一致していない場合は、エラー画面へ
					location.assign('./error.html');
				} else {
					throw new Error('fetchProduct Response Error')
				}
			})
		},
	}
});
