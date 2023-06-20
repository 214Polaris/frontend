<template>
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <meta name="description" content="">
        <meta name="author" content="Ansonika">
        <title>我的购物车</title>
    </head>

    <body>
        <div id="page">

            <main class="bg_gray">
                <div class="container margin_30">
                    <div class="page_header">
                        <h1>我的购物车</h1>
                    </div>
                    <!-- /page_header -->
                    <table class="table table-striped cart-list">
                        <thead>
                            <tr>
                                <th>
                                    商品
                                </th>
                                <th>
                                    单价
                                </th>
                                <th>
                                    数量
                                </th>
                                <th>
                                    总价
                                </th>
                                <th>

                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <div v-for="(item,index) in cartlist" :key="index">
                            <tr>
                                <td>
                                    <div class="thumb_cart">
                                        <img :src="item.image.URL"
                                            data-src="item.image.URL" class="lazy" alt="Image">
                                    </div>
                                    <span class="item_cart">{{item.productName}}</span>
                                </td>
                                <td>
                                    <strong>{{productPrice}}</strong>
                                </td>
                                <td>
                                    <div class="numbers-row">
                                        <input type="text" value="1" id="quantity_1" class="qty2" name="quantity_1">
                                        <div class="inc button_inc">+</div>
                                        <div class="dec button_inc">-</div>
                                    </div>
                                </td>
                                <td>
                                    <strong>$140.00</strong>
                                </td>
                                <td class="options">
                                    <a href="#"><i class="ti-trash"></i></a>
                                </td>
                            </tr>
                        </div>
                            <tr>
                                <td>
                                    <div class="thumb_cart">
                                        <img src="img/products/product_placeholder_square_small.jpg"
                                            data-src="img/products/shoes/2.jpg" class="lazy" alt="Image">
                                    </div>
                                    <span class="item_cart">Armor Okwahn II</span>
                                </td>
                                <td>
                                    <strong>$110.00</strong>
                                </td>
                                <td>
                                    <div class="numbers-row">
                                        <input type="text" value="1" id="quantity_2" class="qty2" name="quantity_2">
                                        <div class="inc button_inc">+</div>
                                        <div class="dec button_inc">-</div>
                                    </div>
                                </td>
                                <td>
                                    <strong>$110.00</strong>
                                </td>
                                <td class="options">
                                    <a href="#"><i class="ti-trash"></i></a>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <div class="thumb_cart">
                                        <img src="img/products/product_placeholder_square_small.jpg"
                                            data-src="img/products/shoes/3.jpg" class="lazy" alt="Image">
                                    </div>
                                    <span class="item_cart">Armor Air Wildwood ACG</span>
                                </td>
                                <td>
                                    <strong>$90.00</strong>
                                </td>

                                <td>
                                    <div class="numbers-row">
                                        <input type="text" value="1" id="quantity_3" class="qty2" name="quantity_3">
                                        <div class="inc button_inc">+</div>
                                        <div class="dec button_inc">-</div>
                                    </div>
                                </td>
                                <td>
                                    <strong>$90.00</strong>
                                </td>
                                <td class="options">
                                    <a href="#"><i class="ti-trash"></i></a>
                                </td>
                            </tr>
                        </tbody>
                    </table>

                    <div class="row add_top_30 flex-sm-row-reverse cart_actions">
                        <div class="col-sm-4 text-end">
                            <button type="button" class="btn_1 gray">Update Cart</button>
                        </div>
                        <div class="col-sm-8">
                            <div class="apply-coupon">
                                <div class="form-group">
                                    <div class="row g-2">
                                        <div class="col-md-6"><input type="text" name="coupon-code" value=""
                                                placeholder="Promo code" class="form-control"></div>
                                        <div class="col-md-4"><button type="button" class="btn_1 outline">Apply
                                                Coupon</button></div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <!-- /cart_actions -->

                </div>
                <!-- /container -->

                <div class="box_cart">
                    <div class="container">
                        <div class="row justify-content-end">
                            <div class="col-xl-4 col-lg-4 col-md-6">
                                <ul>
                                    <li>
                                        <span>Total</span> $247.00
                                    </li>
                                </ul>
                                <a href="cart-2.html" class="btn_1 full-width cart">Proceed to Checkout</a>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- /box_cart -->

            </main>
            <!--/main-->
        </div>
    </body>
</template>

<script setup>
import { onMounted } from 'vue';
let cartlist = []
//获取购物车内容
onMounted(() => {
    axios({
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        method: 'get',
        url: '/api/cart',
    })
        .then((response) => {
            cartlist = response.data;
            console.log(cartlist)
        })
        .catch((error) => {
            console.error(error);
        });
});
</script>