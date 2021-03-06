const form = $("#user-input");
const list = $("#conversation_list");

form.keypress(function(event){
    if(event.keyCode != 13){ // ENTER
        return;
    }

    event.preventDefault(); // don't refresh the page.
    const userText = form.val(); // get the text from the input form
    form.val(" "); // wipes the text box.
    
    // add input to the list
    list.append('<li class="list-group-item list-group-item-success text-right">' +"User : " + userText + "</li>");

    // GET/POST
    const queryParams = {"user-input" : userText }
    $.get("/chat", queryParams)
        .done(function(resp){
            const newItem = '<li  class="list-group-item list-group-item-danger">'+"ELiza : " + resp + "</li>";
            setTimeout(function(){
                list.append(newItem)
                // for the auto fix to the bottom go to https://stackoverflow.com/questions/47425453/html-css-auto-scroll-page-to-bottom
                $("html, body").scrollTop($("body").height());
            }, 1000);//set timeout to give wait to response
            
        }).fail(function(){
            const newItem = "<li class='list-group-item list-group-item-danger' >Come back Later.</li>";
            list.append(newItem);
            // window.scrollTo(0,document.body.scrollHeight); //scroll to the bottom so the latest chat is in view
        });
       
        
});