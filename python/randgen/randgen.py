import random


class RandomGen(object):
    """"
    RandomGen represents a random number generator with fixed random number set and
    probability distribution. Once called the generator can be used to generate numbers
    from the set according to the probability distribution. 
    """
    # Values that may be returned by next_num()
    _random_nums = [-1, 0, 1, 2, 3]
    # Probability of the occurrence of random_nums
    _probabilities = [0.01, 0.3, 0.58, 0.1, 0.01]

    # cumulative probability domain
    _domain = []

    @property
    def domain(self):
        return self._domain

    def __init__(self) -> None:
        # create cumulative probability domain
        self._domain = self._calculate_cumulative_probabilities()

    def _calculate_cumulative_probabilities(self):
        return [round(sum(self._probabilities[:i+1]), 2) for i in range(len(self._probabilities))]

            
    def next_num(self):
        """
        Returns one of the randomNums. When this method is called multiple
        times over a long period, it should return the numbers roughly with
        the initialized probabilities
        """
        rnd = random.random()
        for num, prob in zip(self._random_nums, self._domain):
            if rnd < prob:
                return num

