new Vue({
	// 「el」プロパティーで、Vueの表示を反映する場所=HTML要素のセレクター（id）を定義
	el: '#app',

	// data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
	data: {
		// 掲示板情報
		products: [],
		// 宛先
		productName: '',
		// メッセージ
		productMessage: '',
		// パスワード
		productPassword: '',
		// true:入力済・false:未入力
		isEntered: false
	},

	// 算出プロパティ
	computed: {
		// 掲示板情報の一覧を表示する
		labels() {
			return this.options.reduce(function (a, b) {
				return Object.assign(a, { [b.value]: b.label })
			}, {})
		},
		// 表示対象の掲示板情報を返却する
		computedProducts() {
			return this.products.filter(function (el) {
				var option = true
				return option
			}, this)
		},
		// 入力チェック
		validate() {
			var isEnteredProductName = 0 < this.productName.length
			this.isEntered = isEnteredProductName
			return isEnteredProductName
		}
	},

	// インスタンス作成時の処理
	created: function() {
		this.doFetchAllProducts()
	},

	// メソッド定義
	methods: {
		// 全ての掲示板情報を取得する
		doFetchAllProducts() {
			axios.get('/fetchAllProducts')
			.then(response => {
				if (response.status != 200) {
					throw new Error('fetchAllProducts Response Error')
				} else {
					var resultProducts = response.data

					// サーバから取得した掲示板情報をdataに設定する
					this.products = resultProducts
				}
			})
		},
		// 1つの掲示板情報を取得する
		doFetchProduct(ID, Password) {
			axios.get('/fetchProduct', {
				params: {
					productID: ID,
					productPassword: Password
				}
			})
			.then(response => {
				if (response.status != 200) {
					throw new Error('fetchProduct Response Error')
				} else {
					var resultProducts = response.data

					// 選択された掲示板情報のインデックスを取得する
//					var index = this.products.indexOf(product)
					var index = ID

					// spliceを使うとdataプロパティの配列の要素をリアクティブに変更できる
					this.products.splice(index, 1, resultProducts[0])
				}
			})
		},
		// 掲示板情報を登録する
		doAddProduct() {
			// サーバへ送信するパラメータ
			const params = new URLSearchParams()
			params.append('productName', this.productName)
			params.append('productMessage', this.productMessage)

			axios.post('/addProduct', params)
			.then(response => {
				if (response.status != 200) {
					throw new Error('addProduct Response Error')
				} else {
					// 掲示板情報を取得する
					this.doFetchAllProducts()

					// 入力値を初期化する
					this.initInputValue()
				}
			})
		},
		// 掲示板情報を削除する
		doDeleteProduct(ID) {
			// サーバへ送信するパラメータ
//			const params = new URLSearchParams()
//			params.append('productID', product.ID)
console.log(ID)

			axios.post('/deleteProduct', params)
			.then(response => {
				if (response.status != 200) {
					throw new Error('deleteProduct Response Error')
				} else {
					// 掲示板情報を取得する
					this.doFetchAllProducts()
				}
			})
		},
		// パスワード処理
		openPasswordPage(item) {
console.debug(item);
			let password = prompt('パスワードを入力してください');
console.log(password);
console.log(item.Password);
			if ( password === item.Password) {
				console.log('パスワードは一致しています');
				console.log(item.Name);
				console.log(item.Message);
				let baseurl = './message.html';

				let urlParameter = {
					name: item.Name,
					message: item.Message
				};
				let newurl = baseurl + "?" + 
					Object.entries(urlParameter).map((e) => {
							let key = e[0];
							let value = encodeURI(e[1]);
							return `${key}=${value}`;
						}).join("&");
				console.log(newurl);
				location.href = newurl;
			} else {
				console.log('パスワードが違います');
				location.assign('./error.html');
			}
		},
		// パスワードを照合する
		collationPassword(ID) {i
console.log(ID);
		},
		// 入力値を初期化する
		initInputValue() {
			this.productName = ''
			this.productMessage = ''
		}
	}
})
