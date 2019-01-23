function loadEditor() {
//	const save = document.getElementById('save')
	let editor = ace.edit("editor");
	editor.setTheme("ace/theme/monokai");
	editor.setFontSize(14);
	editor.getSession().setMode("ace/mode/html");
	editor.getSession().setUseWrapMode(true); //折り返し
	editor.getSession().setTabSize(2); //タブ幅
//	save.addEventListener("click", () => {
//		saveFile(editor)
//	})
	editor.addEventListener("change", () => {console.log(editor.getValue())})
}

window.addEventListener("load", () => loadEditor())


function saveFile(editor) {
	console.log(editor)
}
