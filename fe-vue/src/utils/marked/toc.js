export default class MarkedToc {
	constructor() {
		this.toc = [];
		this.index = 0;
		this.result = "";
	}
	add(text, level) {
		const anchor = `#toc${level}${this.index - 0}`;
		this.toc.push({
			anchor,
			level,
			text
		});
		this.index += 1;
		return anchor;
	}
	toHTML() {
		let levelStack = [];
		this.toc.forEach(item => {
			let levelIndex = levelStack.indexOf(item.level);
			// 没有找到相应level的ul标签，则将li放入新增的ul中
			if (levelIndex === -1) {
				levelStack.unshift(item.level);
				this.addStartUL();
				this.addLI(item);
			} // 找到了相应level的ul标签，并且在栈顶的位置则直接将li放在此ul下
			else if (levelIndex === 0) {
				this.addLI(item);
			} // 找到了相应level的ul标签，但是不在栈顶位置，需要将之前的所有level出栈并且打上闭合标签，最后新增li
			else {
				while (levelIndex--) {
					levelStack.shift();
					this.addEndUL();
				}
				this.addLI(item);
			}
		});
		// 如果栈中还有level，全部出栈打上闭合标签
		while (levelStack.length) {
			levelStack.shift();
			this.addEndUL();
		}
		// 清理先前数据供下次使用
		this.toc = [];
		this.index = 0;
	}
	addStartUL() {
		this.result += "<ul>";
	}
	addEndUL() {
		this.result += "</ul>";
	}
	addLI({anchor, text, level}) {
		level = level || 0
		this.result += `<li style="padding-left:${(level - 1) * 15}px"><a href="#${anchor}">` + text + "<a></li>";
	}
}
