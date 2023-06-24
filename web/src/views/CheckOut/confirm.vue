<template>
	<div id="page">
		<main class="bg_gray">
			<div class="container">
				<div class="row justify-content-center">
					<div class="col-md-5">
						<div id="confirm">
							<div class="icon icon--order-success svg add_bottom_15">
								<svg xmlns="http://www.w3.org/2000/svg" width="72" height="72">
									<g fill="none" stroke="#8EC343" stroke-width="2">
										<circle cx="36" cy="36" r="35"
											style="stroke-dasharray:240px, 240px; stroke-dashoffset: 480px;"></circle>
										<path d="M17.417,37.778l9.93,9.909l25.444-25.393"
											style="stroke-dasharray:50px, 50px; stroke-dashoffset: 0px;"></path>
									</g>
								</svg>
							</div>
							<h2>订单已完成</h2>
							<p>多谢惠顾！请点击下方按钮回到主页</p>
							<el-button class="btn btn-primary" @click="goToHomePage">回到主页</el-button>
						</div>
					</div>
				</div>
				<!-- /row -->
			</div>
			<!-- /container -->
		</main>
	</div>
</template>
<script setup>
import axios from 'axios';
import { onMounted } from 'vue';
onMounted(() => {
	const hashIndex = window.location.href.indexOf('#');
    const queryString = window.location.href.slice(hashIndex + 1);
    const urlParams = new URLSearchParams(queryString);
    const outTradeNo = urlParams.get('out_trade_no');
	axios({
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
		},
		method: "post",
		url: "/api/finishorder",
		data: {
			orderId: outTradeNo
		},
	})
		.then((response) => {
			console.log(response);
		})
		.catch((error) => {
			console.error(error);
		});
}
)
function goToHomePage() {
	window.location.href = "/"; // 替换为您的主页地址
};
</script>