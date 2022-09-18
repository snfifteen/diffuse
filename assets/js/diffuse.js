$(document).ready(function(){
    $("#btnGenerate").click(function(){
        var promptInput = $("#tbxPrompt").val();

        if (promptInput !== ""){
            $.post("/generate",
            {
                prompt: promptInput
            },
            function(data, status){
                alert("Data: " + data + "\nStatus: " + status);
            });
        }
    });
});

