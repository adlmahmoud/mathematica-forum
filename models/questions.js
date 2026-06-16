const mongoose = require("mongoose");

const answerSchema =
new mongoose.Schema({

  author:{
    type:mongoose.Schema.Types.ObjectId,
    ref:"User"
  },

  body:{
    type:String,
    required:true
  },

  accepted:{
    type:Boolean,
    default:false
  },

  votes:{
    type:Number,
    default:0
  }

},
{
  timestamps:true
});

const questionSchema =
new mongoose.Schema({

  title:{
    type:String,
    required:true
  },

  body:{
    type:String,
    required:true
  },

  tags:[
    String
  ],

  author:{
    type:mongoose.Schema.Types.ObjectId,
    ref:"User"
  },

  votes:{
    type:Number,
    default:0
  },

  solved:{
    type:Boolean,
    default:false
  },

  answers:[
    answerSchema
  ]

},
{
  timestamps:true
});

module.exports = mongoose.model(
  "Question",
  questionSchema
);