// Vue
var app = new Vue({
	el: "#app",
	data: {
		productID: '',
		productName: '',
		productMessage: '',
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
			let Id = params.get('id');

			axios.get('/fetchProduct', {
				params: {
					productID: Id
				}
			})
			.then(response => {
				if (response.status != 200) {
					throw new Error('fetchProduct Response Error')
				} else {
					var resultProducts = response.data

					// 掲示板表示タグ生成
					this.productName = resultProducts.Name;
					this.productMessage = resultProducts.Message;
				}
			})
		},
	}
});
