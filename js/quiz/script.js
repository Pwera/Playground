var quizController = (function() {
  function Question(id, questionText, options, correctAnswer) {
    this.id = id;
    this.questionText = questionText;
    this.options = options;
    this.correctAnswer = correctAnswer;
  }

  var questionLocalStorage = {
    setQuestionCollection: function(newCollection) {
      localStorage.setItem("questionCollection", JSON.stringify(newCollection));
    },
    getQuestionCollection: function() {
      return JSON.parse(localStorage.getItem("questionCollection"));
    },
    removeQuestionCollection: function() {
      localStorage.removeItem("questionCollection");
    }
  };

  return {
    addQuestionOnLS: function(newQuestionText, options) {
      var optionsArr = [],
        corrAns,
        questionId = 0,
        newQuestion,
        getStoredQuestions;
      for (var i = 0; i < options.length; i++) {
        if (options[i].value !== "") {
          optionsArr.push(options[i].value);
        }
        if (options[i].previousElementSibling.checked && options[i] !== "") {
          corrAns = options[i].value;
        }
      }
      if (questionLocalStorage.getQuestionCollection().length > 0) {
        questionId =
          questionLocalStorage.getQuestionCollection()[
            questionLocalStorage.getQuestionCollection().length - 1
          ].id + 1;
      } else {
        questionId = 0;
      }
      newQuestion = new Question(
        questionId,
        newQuestionText.value,
        optionsArr,
        corrAns
      );
      getStoredQuestions = questionLocalStorage.getQuestionCollection();
      getStoredQuestions.push(newQuestion);
      questionLocalStorage.setQuestionCollection(getStoredQuestions);
      console.log(questionLocalStorage.getQuestionCollection());
    }
  };
})();

var UIController = (function() {
  var domItems = {
    questionInsertBttn: document.getElementById("question-insert-btn"),
    newQuestionText: document.getElementById("new-question-text"),
    adminOptions: document.querySelectorAll(".admin-option")
  };

  return {
    getDomItems: domItems
  };
})();

var controller = (function(quizCtrl, UICtrl) {
  var selectedDomItems = UICtrl.getDomItems;
  selectedDomItems.questionInsertBttn.addEventListener("click", function() {
    quizCtrl.addQuestionOnLS(
      selectedDomItems.newQuestionText,
      selectedDomItems.adminOptions
    );
  });
})(quizController, UIController);
