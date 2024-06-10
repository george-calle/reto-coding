const matrizSchema = require('../esquema/matriz.model.js');

const calcularMatriz = (req, res) => {
  const { error, value } = matrizSchema.validate(req.body);
  if (error) {
    return res.status(400).json({ error: error.details[0].message });
  }

  const matriz = value.matriz;

  let val_valor_maximo = Number.NEGATIVE_INFINITY;
  let val_valor_minimo = Number.POSITIVE_INFINITY;
  let val_suma_total = 0;
  let val_promedio = 0;
  let val_matriz_diagonal = 0;
  let elementosTotales = 0;

  for (let i = 0; i < matriz.length; i++) {
    for (let j = 0; j < matriz[i].length; j++) {
      const valor = matriz[i][j];
      val_valor_maximo = Math.max(val_valor_maximo, valor);
      val_valor_minimo = Math.min(val_valor_minimo, valor);
      val_suma_total += valor;
      elementosTotales++;
      if (i === j) {
        val_matriz_diagonal += valor;
      }
    }
  }

  val_promedio = val_suma_total / elementosTotales;

  res.json({
    val_valor_maximo,
    val_valor_minimo,
    val_promedio,
    val_suma_total,
    val_matriz_diagonal
  });
};

module.exports = {
  calcularMatriz
};
