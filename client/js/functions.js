var COLORS = [Material.Blue, Material.Purple, Material.Red, Material.Green]

function getMaterialColor() {
    return COLORS[getRandomInt(0, COLORS.length - 1)]
}

function getRandomInt(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min)) + min;
}

function findDiceByFace(face, model) {
    return findModelIndex(model, function(item) {
        return item.face === face
    }) || 0
}

function findModel(model, func) {
    for(var i = 0; i < model.count; ++i) if (func(model.get(i))) return model.get(i)
      return null
}

function findModelIndex(model, func) {
    for(var i = 0; i < model.count; ++i) if (func(model.get(i))) return i
}

function toDateTime(secs) {
    var t = new Date(1970, 0, 1); // Epoch
    t.setSeconds(secs);
    return t.getMinutes() +":"+t.getSeconds();
}
