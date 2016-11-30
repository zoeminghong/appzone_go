/**
 * Created by gJason on 2016/11/30.
 */
$(function () {
    $.ajax({
        url:"/getjson",
        success:function (data) {
            $("#json").text(data)
        }
    })
})