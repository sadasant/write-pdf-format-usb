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
    var alive_interval = setInterval(function() {
        $.ajax({
            method: "GET",
            url: "/keep-alive",
            error: function() {
                clearInterval(alive_interval);
                window.close();
            }
        });
    }, 1000);

    var $content = $("#content");

    function loadView(name) {
        console.log("LOADING", name)
        $.ajax({
            method:  "GET",
            url:     "/views/"+name+".html?no-cache=true", // no-cache is just to avoid browser cache
            success: function(data) {
                $content.html(data);
            }
        })
    }

    loadView($content.attr("view"));
    $content.attr("view", "");

    // Move to next view
    $(window).on("click", function(e) {
        var $target = $(e.target);
        var view    = $target.attr("view");
        console.log($target, view);
        if (view) loadView(view);
    });

    // Show content based on selects
	$content.on("change", "select", function(e){
		var $this = $(this);
		var show  = $this.attr("show");
		var value = $this.val();
		console.log("Select changed", show, value);
		if (show) {
			if (show[0] === "#") $("[id^="+show.substr(1)+"]").addClass("hidden");
			if (show[0] === ".") $("[class^="+show.substr(1)+"]").addClass("hidden");
			$(show+"-"+value).removeClass("hidden");
		}
	});
})();
