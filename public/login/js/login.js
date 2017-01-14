function SendForm() {
	var email = $('#email').val();
	var password = $('#password').val();
	$.post('/api/user/login', {'email': email,'pass':password}, function(response){
		$('.notification').css('display','none');
		switch (response){
			case "OK_LOGGED_IN":
				$('#ok-login').css('display','block');
				setTimeout(function(){
					// window.location.pathname = "/dashboard"
				},1000)
			case "ERR_WRONG_CREDENTIALS":
				$('#wrong-cred').css('display','block');
			case "ERR_INTERNAL":
				$('#inter-err').css('display','block');
			default:
				$('#inter-err').css('display','block');
		}
		setTimeout(function(){
			$('.notification').css('display','none');
		},2000)
	});
}

$('#login').on('click',SendForm);