function SendForm() {
	var email = $('#email').val();
	var user = $('#user').val();
	var pass = $('#pass').val();

	$.post('/api/user/register', {'email': email,'user':user,'pass':pass}, function(response){
		$('.notification').css('display','none');
		switch(response){
			case "OK_USER_CREATED":
				$('#ok-created').css('display','block');
				setTimeout(function(){
					// window.location.pathname = "/login"
				},1000)
				break;
			case "ERR_USER_TAKEN":
				$('#user-taken').css('display','block');
				break;
			case "ERR_FIELDS_MISSING":
				$('#invalid-form').css('display','block');
				break;
			default:
				$('#inter-err').css('display','block');
		}
		setTimeout(function(){
			$('.notification').css('display','none');
		},2000)
	});
}

$('#register').on('click',SendForm);