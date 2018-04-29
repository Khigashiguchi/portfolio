(function(Vue) {
    "use strict";

    new Vue({
        el: 'body',

        data: {
            blogs: [],
        },

        created: function() {
            this.$http.get('/blogs').then(function(res) {
                this.blogs = res.data.items ? res.data.items : [];
            })
        }
    });
})(Vue);