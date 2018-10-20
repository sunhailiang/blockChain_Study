$("#logout").click(function () {
    $("#cover").css({"display": "block"})
})
$("#cancle").click(function () {
    $("#cover").css({"display": "none"})
})
$("#ok").click(function () {
    $.get("/logout", function (res) {
        if (eval(res) == 302) {
            window.location.href = "/login"
        }
    })
})

$("#select").change(function () {
    var type = $(this).val();
    $.get("/", {"selectType": type}, function (res) {
        console.log(res)
    })
})