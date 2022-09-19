$(document).ready(function(){

    $("#imgWidth").change(function(){
        $("#lblImgWidth").text($(this).val());
    });

    $("#btnGenerate").click(function(){
        var promptInput = $("#tbxPrompt").val();

        if (promptInput !== ""){
            $.post("/generate",
            {
                prompt: promptInput,
                width: $("#imgWidth").val(),
                height:$("#imgHeight").val()
            },
            function(data, status){
                alert("Data: " + data + "\nStatus: " + status);
            });
        }
    });
});

