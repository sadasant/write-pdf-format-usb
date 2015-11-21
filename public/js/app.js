(function() {
    // // /exit example
    // setTimeout(function() {
    //     alert("Bye bye!");
    //     $.get("/exit");
    //     setTimeout(function() {
    //         window.close();
    //     }, 250)
    // }, 10*1000)

    // Keep alive
    setInterval(function() {
        $.ajax({
            method: "GET",
            url: "/keep-alive",
            error: function() {
                window.close();
            }
        });
    }, 2000);

    var $content = $("#content");

    function loadView(name) {
        console.log("LOADING", name)
        $.ajax({
            method:  "GET",
            url:     "/views/"+name+".html",
            success: function(data) {
                $content.html(data);
            }
        })
    }

    loadView($content.attr("view"));
    $content.attr("view", "");

    $(window).on("click", function(e) {
        var $target = $(e.target);
        var view    = $target.attr("view");
        console.log($target, view);
        if (view) loadView(view);
    });
})();
