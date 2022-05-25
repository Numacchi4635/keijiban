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
		// 環境変数情報
		pulbic_mode: '',
		// true:入力済・false:未入力
		isEntered: false,
		// タイトル
		title: '',
		// true:掲示板表示・false：掲示板非表示
		isKeijibanDisplay: false,
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
			return this.products.Products
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
		// 掲示板一覧表示
		this.doFetchAllProducts()
	},

	// メソッド定義
	methods: {
		// 全ての掲示板情報を取得する
		doFetchAllProducts() {
			this.superUserPassword = prompt('管理者専用パスワードを入力してください');

			// パスワード未入力時はエラーページへ
			if (this.superUserPassword == null){
				this.isKeijibanDisplay = false
				let url = './superusererror.html';
				location.href = url;
			}

			axios.get('/fetchAllProducts', {
				params: {
					productPassword: this.superUserPassword
				}
			})
			.then(response => {
				if (response.status == 200){
					this.isKeijibanDisplay = true
					var resultProducts = response.data

					// サーバから取得した掲示板情報をdataに設定する
					this.products = resultProducts

					// 取得した環境変数ごとに、タイトルを変更
					if (resultProducts.PublicMode == 'public'){
						this.title = 'パスワード認証付き掲示板'
					} else {
						this.title = '🐹🍎ゆゆこ🍎🐹ファミリーボード返信掲示板'
					}
				}
			})
			.catch(function(error){
				if (error.response.status == 401){
					// パスワードが一致していない場合はエラーページへ
					this.isKeijibanDisplay = false
					let url = './superusererror.html';
					location.href = url;
				} else {
					throw new Error('fetchAllProducts Response Error')
				} 
			})
		},
		// 掲示板情報を削除する
		doDeleteProduct(ID) {
			// 管理者用パスワード入力
			this.superUserPassword  = prompt('管理者専用パスワードを入力してください');

			if ( this.superUserPassword != null ){

				// サーバへ送信するパラメータ
				const params = new URLSearchParams()
				params.append('productID', ID)
				params.append('superUserPassword', this.superUserPassword)

				axios.post('/deleteProduct', params).then(response => {
					if (response.status != 204) {
						throw new Error('deleteProduct Response Error')
					} else {
						// 削除成功時は、index.htmlを更新する
						location.reload();
					}
				})
				.catch(function(error){
					if (error.response.status == 401){
						// パスワードが一致していない場合はエラーページへ
						let url = './superusererror.html';
						location.href = url;
					} else {
						// 上記以外のエラーの場合
						throw new Error('deleteProduct Response Error')
					}
				})
			}
		},
		// メッセージ表示ページへ移動する処理
		openMessagePage(ID) {
			// 一致している場合は、メッセージ表示画面へ
			let baseUrl = './message.html';

			// パラメータ付きURL作成
			let urlParameter = {
				id: ID
			};
			let newUrl = baseUrl + "?" + 
				Object.entries(urlParameter).map((e) => {
					let key = e[0];
					let value = encodeURI(e[1]);
					return `${key}=${value}`;
				}).join("&");
			location.href = newUrl;

		},
		// 入力値を初期化する
		initInputValue() {
			this.productName = ''
			this.productMessage = ''
		}
	}
})
