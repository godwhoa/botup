export default{
	chunk:(m_array,N)=>{
		let chunked = []
		let i,j,temparray,chunk = N;
		for (i=0,j=m_array.length; i<j; i+=chunk) {
		    temparray = m_array.slice(i,i+chunk);
		    chunked.push(temparray)
		}
		return chunked
	}
}