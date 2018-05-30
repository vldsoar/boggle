function resetInputWord(input, grid, stackWord) {
    input.text = ""
    for(var i = 0; i < grid.count; i++) {
        grid.currentIndex = i
        grid.currentItem.state = ""
        stackWord.pop()
    }
    grid.currentIndex = 0
}

function pushWord(word, list) {
    list.model.append({word: word})
}
