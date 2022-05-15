// Vue
var app = new Vue({
	el: "#app",
		data: {
			productID: '',
			productName: '',
			productMessage: '',
		},
	methods: {
		window:onload = function(){

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
console.log(this.productName)
console.log(this.productMessage)
//					let displayTag = '<table><thead><tr><th class="name">宛先</th><th class="message">メッセージ</th></tr></thead><tbody><tr><td class="name">' + resultProducts.Name + '</td>' + '<td class="message">' + resultProducts.Message + '</td></tr></tbody></table>';

					// メッセージ表示
//					let element = document.getElementById('messageinfo');
//					element.insertAdjacentHTML('afterend',  displayTag);
				}
			})
		},
	}
});
