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
})();
