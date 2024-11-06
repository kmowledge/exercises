# it's a game, where you have to guess a 4-digit code. For the correctly guessed numbers on the right position you receive 1 Spirit
# and if it's in wrong position, but good number it's 1 Gnome.

from random import randint


def spirits_and_gnomes():
      secret = list()
      [secret.append(randint(0,9)) for d in range(4)]
      
      guess = input(f"Guess the secret 4-digit code.\nFor each guess you collect a gnome, or spirit if it's in the right position in code.\n")
      while not(len(guess) == 4 and guess.isdigit()):
            guess = input(f'Once again, 4 digits please.\n')
      guess = [int(val) for val in list(guess)]

      spirits = int()
      gnomes = int()

      for i, d in enumerate(secret):
            if secret[i] == guess[i]:
                  secret.pop(i)
                  guess.pop(i)
                  spirits += 1
      
      # gnomes = [gnomes + 1 for val in guess if val in secret] //list comprehensions return type 'list', not int
      # for val in guess: if val in secret: gnomes += 1 //can't write for+if without block with intends
      gnomes = sum(1 for val in guess if val in secret)

      print(str(spirits) + 'S' + str(gnomes) + 'G')
      print('The secret number was: ' + ''.join(map(str, secret)) + '.')




spirits_and_gnomes()











# errors to investigate
#       for (guess := input(f"Guess the secret 4-digit code.\n"), guess == False, guess++):
#       [guess = input(f'once again, 4 digits please\n') for not(len(guess) == 4 & type(guess) == int)] nie można przypisywać w 1. arg, ani 
