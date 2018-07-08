function search() {
    //Get "searching for"
    var name = document.getElementById("search-form").elements["search-val"].value;
    var pattern = name.toLowerCase();
    var targetId = "";

    //Search in divs
    var divs = document.getElementsByClassName("search-item");

    for (var i = 0; i < divs.length; i++) {
        var para = divs[i].getElementsByTagName("p");
        var index = para[0].innerText.toLowerCase().indexOf(pattern);
        if (index != -1) {
            targetId = divs[i].parentNode.id;
            console.log("search")
            document.getElementById(targetId).scrollIntoView();
            break;
        }
    }
};