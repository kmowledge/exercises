// team_score and personal_score are private properties, and getScore() is a privileged method.
// team_score is a shared variable to all the Player instances.

var Player = (function() {
      var team_score = 0; // private property, class-shared variable

      function innerPlayer(name) {
            this.name = name; // public property
            var personal_score = 0; // private property, instance-specific variable
      

            this.score = function() {
                  team_score += 1;
                  personal_score += 1;
            };

            this.getScore = function(arg) {
                  if (arg == 'team') {
                        return team_score;
                  } else if (arg == 'personal') {
                        return `${this.name} scored ${personal_score}.`;
                  } else {
                        return `Personal: ${personal_score} | Team: ${team_score}`;
                  }
            };
      }

      return innerPlayer
}());

player1 = new Player("Tobiasz");
player2 = new Player("Piotr");

console.log(player1.name) // "Tobiasz"
console.log(player1.personal_score) // undefined
console.log(player1.team_score) // undefined

console.log(player1.getScore()); // undefined (no parameter provided)
player1.score();
player2.score();
player2.score();
console.log(player1.getScore('team')); // 3
console.log(player2.getScore('personal')) // 2