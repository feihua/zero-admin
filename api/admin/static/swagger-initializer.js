window.onload = function () {
    //<editor-fold desc="Changeable Configuration Block">

    // the following lines will be replaced by docker/configurator, when it runs in a docker-container
    window.ui = SwaggerUIBundle({
        urls: [
            {name: "系统管理", url: "./sys.json"},
            {name: "订单管理", url: "./oms.json"},
            {name: "商品管理", url: "./pms.json"},
            {name: "会员管理", url: "./ums.json"},
            {name: "营销管理", url: "./sms.json"},
            {name: "内容管理", url: "./cms.json"}
        ],
        dom_id: '#swagger-ui',
        deepLinking: true,
        presets: [
            SwaggerUIBundle.presets.apis,
            SwaggerUIStandalonePreset
        ],
        plugins: [
            SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout"
    });

    //</editor-fold>
};
