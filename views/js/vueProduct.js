new Vue({
	// 「el」プロパティーで、Vueの表示を反映する場所=HTML要素のセレクター（id）を定義
	el: '#app',

	// data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
	data: {
		// 掲示板情報
		products: [],
		// ID
		productID: '',
		// 宛先
		productName: '',
		// メッセージ
		productMessage: '',
		// パスワード
		productPassword: '',
		// SuperUser パスワード
		superUserPassword: '',
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
			const params = new URLSearchParams()
			params.append('productID', ID)

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
		// 管理者専用パスワード処理
		openSuperUserPassword() {
			let password = prompt('管理者専用パスワードを入力してください');

			// サーバーにパスワードが一致しているか問い合わせる
			const params = new URLSearchParams()
			params.append('superUserPassword', password)
			axios.post('/superUserPasswordCollation', params)
			.then(response => {
				if (response.status == 200) {
					// パスワードが一致している場合はmessagecreate.htmlへ
					let url = './messagecreate.html';
					location.href = url;
				} else if ( response.status == 201) {
					// パスワードが一致していない場合はエラーページへ
					let url = './superusererror.html';
					location.href = url;
				} else {
					// 上記以外のエラーの場合
					throw new Error('fetchProduct Response Error')
				}
			})
		},
		// パスワード処理
		openPasswordPage(item) {
console.log(item);

			// パスワード入力ダイアログ表示
			let password = prompt('パスワードを入力してください');

			// パスワードを照合する
			if ( password != null ){	// キャンセルの場合は何もしない
				const params = new URLSearchParams()
				params.append('productID', item.ID)
				params.append('productPassword', password)

				axios.post('/UserPasswordCollation', params)
				.then(response => {
					if ( response.status == 200 ) {

						// 一致している場合は、メッセージ表示画面へ
						let baseurl = './message.html';
						// パラメータ付きURL作成
						let urlParameter = {
							id: item.ID
						};
						let newurl = baseurl + "?" + 
							Object.entries(urlParameter).map((e) => {
									let key = e[0];
									let value = encodeURI(e[1]);
									return `${key}=${value}`;
								}).join("&");
						location.href = newurl;
					} else if (response.status == 201){

						// パスワードが一致していない場合は、エラー画面へ
						location.assign('./error.html');
					} else {
						throw new Error('fetchProduct Response Error')
					}
				})
			}
		},
		// 入力値を初期化する
		initInputValue() {
			this.productName = ''
			this.productMessage = ''
		}
	}
})
