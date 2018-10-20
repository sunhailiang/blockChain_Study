$("#addArticleBtn").click(function () {
    var img_upload = $("#input_file")[0]
    readFile(img_upload)
})

function readFile(img) {
    var file = img.files[0];
    if (!/image\/\w+/.test(file.type)) {
        alert("请确保文件为图像类型");
        return false;
    }
    if (file.size > 2000000) {
        alert("图片过大");
        return false;
    }
    console.log(file)
    var reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = function () {
        var articleName = $("#articleName").val()
        var Type = $("#sel_opt").val()
        var content = $("#input_multxt").val()
        var blobData = dataURLtoBlob(reader.result)
        $.post("/addarticle", {
            "articleName": articleName,
            "Type": Type,
            "content": content,
            "img": blobData
        }, function (res) {
            console.log(res)
        })
    }
}

function dataURLtoBlob(dataurl) {
    var arr = dataurl.split(','), mime = arr[0].match(/:(.*?);/)[1],
        bstr = atob(arr[1]), n = bstr.length, u8arr = new Uint8Array(n);
    while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
    }
    return new Blob([u8arr], {type: mime});
}
