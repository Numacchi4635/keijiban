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

			// URLパラメータ取得
			let url = new URL(window.location.href);
			let params = url.searchParams;
			let password = params.get('pass');
			let name = params.get('name');
			if (name == ''){
				// URLパラメータのnameがnullならばエラー画面へ
				location.assign('./errornoname.html');
			}

			axios.get('/fetchProduct', {
				params: {
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
