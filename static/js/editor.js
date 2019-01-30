function roadEditor() {
	const save = document.getElementById('save_button')
	var connection = new WebSocket("ws://127.0.0.1:8080/room");
	let editor = ace.edit("editor");
	editor.setTheme("ace/theme/monokai");
	editor.setFontSize(14);
	editor.getSession().setMode("ace/mode/html");
	editor.getSession().setUseWrapMode(true); //折り返し
	editor.getSession().setTabSize(2); //タブ幅
	save.addEventListener("click", () => {
		saveFile(editor)
	})
	editor.addEventListener("change", () => {
		connection.send(editor.getValue());
	})
	connection.onmessage = function(e) {
		if ( editor.getValue() != e.data ) {
			editor.setValue(e.data);
		}
	}
}

function saveFile(editor){
	console.log(editor.getValue())
}
window.addEventListener("load", () => roadEditor())
