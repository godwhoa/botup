export default{
	show:(selector, display)=>{
		display = display || 'block'
		let elements = document.querySelectorAll(selector)
		for (let i = 0; i < elements.length; i++) {
			elements[i].style.display = display
		}
	},
	hide:(selector)=>{
		let elements = document.querySelectorAll(selector)
		for (let i = 0; i < elements.length; i++) {
			elements[i].style.display = 'none'
		}
	}
}