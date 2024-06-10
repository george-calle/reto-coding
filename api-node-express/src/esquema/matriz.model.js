const Joi = require('joi');

const matrizSchema = Joi.object({
  matriz: Joi.array().items(
    Joi.array().items(
      Joi.number().required()
    ).required()
  ).required()
});

module.exports = matrizSchema;
