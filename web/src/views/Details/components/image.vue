<template>
  <head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, shrink-to-fit=no"
    />
    <title>商品详情</title>
  </head>

  <body>
    <!--主体部分-->
    <main class="bg_gray">
      <div class="container margin_30">
        <div class="page_header">
          <h1>Armor Air X Fear</h1>
        </div>
        <!-- /page_header -->
        <div class="row justify-content-center">
          <div class="col-lg-8">
            <div class="owl-carousel owl-theme prod_pics magnific-gallery">
              <!--商品图片栏1-->
              <div class="item">
                <a
                  href="../../../../public/img/products/shoes/product_detail_1.jpg"
                  title="Photo title"
                  data-effect="mfp-zoom-in"
                  ><img src="../../../../public/img/products/shoes/product_detail_1.jpg" alt=""
                /></a>
              </div>
              <!--商品图片栏2-->
              <div class="item">
                <a
                  href="../../../../public/img/products/shoes/product_detail_2.jpg"
                  title="Photo title"
                  data-effect="mfp-zoom-in"
                  ><img
                    src="../../../../public/img/products/product_placeholder_detail_2.jpg"
                    data-src="../../../../public/img/products/shoes/product_detail_2.jpg"
                    alt=""
                    class="owl-lazy"
                /></a>
              </div>
              <!-- /item -->
            </div>
            <!-- /carousel -->
          </div>
        </div>
        <!-- /row -->
      </div>
      <!-- /container -->
      <!--主界面部分-->
      <div class="bg_white" style="text-align: center">
        <div class="container">
          <div class="col-lg-12">
            <div class="prod_options version_2">
              <div class="row">
                <label
                  class="col-xl-7 col-lg-5 col-md-6 col-6 pt-0"
                  style="font-size: large"
                  ><strong>颜色</strong></label
                >
                <div class="col-xl-5 col-lg-5 col-md-6 col-6 colors">
                  <ul>
                    <li><a href="#0" class="color color_1 active"></a></li>
                    <li><a href="#0" class="color color_2"></a></li>
                    <li><a href="#0" class="color color_3"></a></li>
                    <li><a href="#0" class="color color_4"></a></li>
                  </ul>
                </div>
              </div>
              <div class="row">
                <label
                  class="col-xl-7 col-lg-5 col-md-6 col-6"
                  style="font-size: large"
                  ><strong>大小</strong> - 大小选择
                  <a
                    href="#0"
                    data-bs-toggle="modal"
                    data-bs-target="#size-modal"
                    ><i class="ti-help-alt"></i></a
                ></label>
                <div class="col-xl-5 col-lg-5 col-md-6 col-6">
                  <div class="custom-select-form">
                    <select class="wide">
                      <option value="" @selected="">Small (S)</option>
                      <option value="">M</option>
                      <option value=" ">L</option>
                      <option value=" ">XL</option>
                    </select>
                  </div>
                </div>
              </div>
              <div class="row">
                <label
                  class="col-xl-7 col-lg-5 col-md-6 col-6"
                  style="font-size: large"
                  ><strong>数量</strong></label
                >
                <div class="col-xl-5 col-lg-5 col-md-6 col-6">
                  <div class="numbers-row">
                    <input
                      type="text"
                      value="1"
                      id="quantity_1"
                      class="qty2"
                      name="quantity_1"
                    />
                    <div class="inc button_inc">+</div>
                    <div class="dec button_inc">-</div>
                  </div>
                </div>
              </div>
              <div class="row mt-3">
                <div class="col-lg-7 col-md-6">
                  <div class="price_main">
                    <span class="new_price">$148.00</span
                    ><span class="percentage">-20%</span>
                    <span class="old_price">$160.00</span>
                  </div>
                </div>
                <div class="col-lg-5 col-md-6">
                  <div class="btn_add_to_cart">
                    <a href="#0" class="btn_1">添加到购物车</a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <!-- /row -->
      </div>
      <!-- /bg_white -->
      <div class="tabs_product bg_white version_2">
        <div class="container">
          <ul class="nav nav-tabs" role="tablist">
            <li class="nav-item">
              <a
                id="tab-A"
                href="#pane-A"
                class="nav-link active"
                data-bs-toggle="tab"
                role="tab"
                >Description</a
              >
            </li>
            <li class="nav-item">
              <a
                id="tab-B"
                href="#pane-B"
                class="nav-link"
                data-bs-toggle="tab"
                role="tab"
                >Reviews</a
              >
            </li>
          </ul>
        </div>
      </div>
      <!-- 简述与评论模块 -->
      <div class="tab_content_wrapper">
        <div class="container">
          <div class="tab-content" role="tablist">
            <div
              id="pane-A"
              class="card tab-pane fade active show"
              role="tabpanel"
              aria-labelledby="tab-A"
            >
              <div class="card-header" role="tab" id="heading-A">
                <h5 class="mb-0">
                  <a
                    class="collapsed"
                    data-bs-toggle="collapse"
                    href="#collapse-A"
                    aria-expanded="false"
                    aria-controls="collapse-A"
                  >
                    Description
                  </a>
                </h5>
              </div>
              <div
                id="collapse-A"
                class="collapse"
                role="tabpanel"
                aria-labelledby="heading-A"
              >
                <div class="card-body">
                  <div class="row justify-content-between">
                    <div class="col-lg-6">
                      <h3>Details</h3>
                      <p>
                        Lorem ipsum dolor sit amet, in eleifend
                        <strong>inimicus elaboraret</strong> his, harum
                        efficiendi mel ne. Sale percipit vituperata ex mel, sea
                        ne essent aeterno sanctus, nam ea laoreet civibus
                        electram. Ea vis eius explicari. Quot iuvaret ad has.
                      </p>
                      <p>
                        Vis ei ipsum conclusionemque. Te enim suscipit recusabo
                        mea, ne vis mazim aliquando, everti insolens at sit. Cu
                        vel modo unum quaestio, in vide dicta has. Ut his laudem
                        explicari adversarium, nisl
                        <strong>laboramus hendrerit</strong> te his, alia
                        lobortis vis ea.
                      </p>
                      <p>
                        Perfecto eleifend sea no, cu audire voluptatibus eam. An
                        alii praesent sit, nobis numquam principes ea eos, cu
                        autem constituto suscipiantur eam. Ex graeci elaboraret
                        pro. Mei te omnis tantas, nobis viderer vivendo ex has.
                      </p>
                    </div>
                    <div class="col-lg-5">
                      <h3>Specifications</h3>
                      <div class="table-responsive">
                        <table class="table table-sm table-striped">
                          <tbody>
                            <tr>
                              <td><strong>Color</strong></td>
                              <td>Blue, Purple</td>
                            </tr>
                            <tr>
                              <td><strong>Size</strong></td>
                              <td>150x100x100</td>
                            </tr>
                            <tr>
                              <td><strong>Weight</strong></td>
                              <td>0.6kg</td>
                            </tr>
                            <tr>
                              <td><strong>Manifacturer</strong></td>
                              <td>Manifacturer</td>
                            </tr>
                          </tbody>
                        </table>
                      </div>
                      <!-- /table-responsive -->
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <!-- /TAB A -->
            <div
              id="pane-B"
              class="card tab-pane fade"
              role="tabpanel"
              aria-labelledby="tab-B"
            >
              <div class="card-header" role="tab" id="heading-B">
                <h5 class="mb-0">
                  <a
                    class="collapsed"
                    data-bs-toggle="collapse"
                    href="#collapse-B"
                    aria-expanded="false"
                    aria-controls="collapse-B"
                  >
                    Reviews
                  </a>
                </h5>
              </div>
              <div
                id="collapse-B"
                class="collapse"
                role="tabpanel"
                aria-labelledby="heading-B"
              >
                <div class="card-body">
                  <div class="row justify-content-between">
                    <div class="col-lg-5">
                      <div class="review_content">
                        <div class="clearfix add_bottom_10">
                          <span class="rating"
                            ><i class="icon-star"></i><i class="icon-star"></i
                            ><i class="icon-star"></i><i class="icon-star"></i
                            ><i class="icon-star"></i><em>5.0/5.0</em></span
                          >
                          <em>Published 54 minutes ago</em>
                        </div>
                        <h4>"Commpletely satisfied"</h4>
                        <p>
                          Eos tollit ancillae ea, lorem consulatu qui ne, eu
                          eros eirmod scaevola sea. Et nec tantas accusamus
                          salutatus, sit commodo veritus te, erat legere fabulas
                          has ut. Rebum laudem cum ea, ius essent fuisset ut.
                          Viderer petentium cu his.
                        </p>
                      </div>
                    </div>
                    <div class="col-lg-5">
                      <div class="review_content">
                        <div class="clearfix add_bottom_10">
                          <span class="rating"
                            ><i class="icon-star"></i><i class="icon-star"></i
                            ><i class="icon-star"></i
                            ><i class="icon-star empty"></i
                            ><i class="icon-star empty"></i
                            ><em>4.0/5.0</em></span
                          >
                          <em>Published 1 day ago</em>
                        </div>
                        <h4>"Always the best"</h4>
                        <p>
                          Et nec tantas accusamus salutatus, sit commodo veritus
                          te, erat legere fabulas has ut. Rebum laudem cum ea,
                          ius essent fuisset ut. Viderer petentium cu his.
                        </p>
                      </div>
                    </div>
                  </div>
                  <!-- /row -->
                  <div class="row justify-content-between">
                    <div class="col-lg-5">
                      <div class="review_content">
                        <div class="clearfix add_bottom_10">
                          <span class="rating"
                            ><i class="icon-star"></i><i class="icon-star"></i
                            ><i class="icon-star"></i><i class="icon-star"></i
                            ><i class="icon-star empty"></i
                            ><em>4.5/5.0</em></span
                          >
                          <em>Published 3 days ago</em>
                        </div>
                        <h4>"Outstanding"</h4>
                        <p>
                          Eos tollit ancillae ea, lorem consulatu qui ne, eu
                          eros eirmod scaevola sea. Et nec tantas accusamus
                          salutatus, sit commodo veritus te, erat legere fabulas
                          has ut. Rebum laudem cum ea, ius essent fuisset ut.
                          Viderer petentium cu his.
                        </p>
                      </div>
                    </div>
                    <div class="col-lg-5">
                      <div class="review_content">
                        <div class="clearfix add_bottom_10">
                          <span class="rating"
                            ><i class="icon-star"></i><i class="icon-star"></i
                            ><i class="icon-star"></i><i class="icon-star"></i
                            ><i class="icon-star"></i><em>5.0/5.0</em></span
                          >
                          <em>Published 4 days ago</em>
                        </div>
                        <h4>"Excellent"</h4>
                        <p>
                          Sit commodo veritus te, erat legere fabulas has ut.
                          Rebum laudem cum ea, ius essent fuisset ut. Viderer
                          petentium cu his.
                        </p>
                      </div>
                    </div>
                  </div>
                  <!-- /row -->
                  <p class="text-end">
                    <a href="leave-review.html" class="btn_1">Leave a review</a>
                  </p>
                </div>
                <!-- /card-body -->
              </div>
            </div>
            <!-- /tab B -->
          </div>
          <!-- /tab-content -->
        </div>
        <!-- /container -->
      </div>
      <!-- /tab_content_wrapper -->
    </main>
    <!-- /main -->
    <!-- 写码表 -->
    <div
      class="modal fade"
      tabindex="-1"
      role="dialog"
      aria-labelledby="size-modal"
      id="size-modal"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Size guide</h5>
            <button
              type="button"
              class="btn-close"
              data-bs-dismiss="modal"
              aria-label="Close"
            ></button>
          </div>
          <div class="modal-body">
            <p>
              Lorem ipsum dolor sit amet, et velit propriae invenire mea, ad nam
              alia intellegat. Aperiam mediocrem rationibus nec te. Tation
              persecuti accommodare pro te. Vis et augue legere, vel labitur
              habemus ocurreret ex.
            </p>
            <div class="table-responsive">
              <table class="table table-striped table-sm sizes">
                <tbody>
                  <tr>
                    <th scope="row">US Sizes</th>
                    <td>6</td>
                    <td>6,5</td>
                    <td>7</td>
                    <td>7,5</td>
                    <td>8</td>
                    <td>8,5</td>
                    <td>9</td>
                    <td>9,5</td>
                    <td>10</td>
                    <td>10,5</td>
                  </tr>
                  <tr>
                    <th scope="row">Euro Sizes</th>
                    <td>39</td>
                    <td>39</td>
                    <td>40</td>
                    <td>40-41</td>
                    <td>41</td>
                    <td>41-42</td>
                    <td>42</td>
                    <td>42-43</td>
                    <td>43</td>
                    <td>43-44</td>
                  </tr>
                  <tr>
                    <th scope="row">UK Sizes</th>
                    <td>5,5</td>
                    <td>6</td>
                    <td>6,5</td>
                    <td>7</td>
                    <td>7,5</td>
                    <td>8</td>
                    <td>8,5</td>
                    <td>9</td>
                    <td>9,5</td>
                    <td>10</td>
                  </tr>
                  <tr>
                    <th scope="row">Inches</th>
                    <td>9.25"</td>
                    <td>9.5"</td>
                    <td>9.625"</td>
                    <td>9.75"</td>
                    <td>9.9375"</td>
                    <td>10.125"</td>
                    <td>10.25"</td>
                    <td>10.5"</td>
                    <td>10.625"</td>
                    <td>10.75"</td>
                  </tr>
                  <tr>
                    <th scope="row">CM</th>
                    <td>23,5</td>
                    <td>24,1</td>
                    <td>24,4</td>
                    <td>24,8</td>
                    <td>25,4</td>
                    <td>25,7</td>
                    <td>26</td>
                    <td>26,7</td>
                    <td>27</td>
                    <td>27,3</td>
                  </tr>
                </tbody>
              </table>
            </div>
            <!-- /table -->
          </div>
        </div>
      </div>
    </div>
  </body>
</template>
