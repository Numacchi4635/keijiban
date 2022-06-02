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
		// true:入力済・false:未入力
		isEntered: false,
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

	// メソッド定義
	methods: {
		// ユーザーパスワード認証処理
		doUserPasswordCollation(){
			// パスワード入力
			let password = prompt('パスワードを入力してください');
			if ( password != null ){
				// URLパラメータ取得
				let url = new URL(window.location.href);
				let params = url.searchParams;

				axios.get('/userPasswordCollation', {
					params: {
						productPassword: password
					}
				})
				.then(response => {
					if (response.status == 200){
						// データ取得
						var resultProducts = response.data
						var encodeUserName = encodeURIComponent(resultProducts.Name)

						// パスワードが一致している場合は、メッセージ表示画面へ
						let baseUrl = './message.html';

						// パラメータ付きURL作成
						let urlParameter = {
							pass: password,
							name: encodeUserName
						};
						let newUrl = baseUrl + "?" + 
							Object.entries(urlParameter).map((e) => {
								let key = e[0];
								let value = encodeURI(e[1]);
								return `${key}=${value}`;
							}).join("&");
						location.href = newUrl;
					}
				})
				.catch(function(error){
					if (error.response.status == 401) {
						// パスワードが一致していない場合は、エラー画面へ
						location.assign('./error.html');
					} else {
						throw new Error('doUserPasswordCollation Response Error')
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
