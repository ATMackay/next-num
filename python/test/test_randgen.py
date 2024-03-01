import unittest
from unittest.mock import patch
from randgen.randgen import RandomGen

class TestRandomGen(unittest.TestCase):

    def test_initialization(self):
        random_gen = RandomGen()
        # Asserting that the cumulative probability domain is initialized correctly
        self.assertEqual(random_gen._domain, [0.01, 0.31, 0.89, 0.99, 1.0])

    def test_domain_length(self):
        random_gen = RandomGen()
        # Asserting that the length of the cumulative probability domain is equal to
        # the number of probabilities and the number of random numbers
        self.assertEqual(len(random_gen._domain), len(RandomGen._probabilities))
        self.assertEqual(len(random_gen._domain), len(RandomGen._random_nums) )

    def test_next_num(self):
        # Mocking the random.random function to return predictable values
        with patch('random.random', side_effect=[0.009, 0.25, 0.5, 0.9, 0.995]):
            
            # create randgen instance
            random_gen = RandomGen()

            # First call, should return 0 as the random value is between 0 and 0.01
            self.assertEqual(random_gen.next_num(), -1)

            # Second call, should return 1 as the random value is between 0.01 and 0.31
            self.assertEqual(random_gen.next_num(), 0)

            # Second call, should return 1 as the random value is between 0.31 and 0.89
            self.assertEqual(random_gen.next_num(), 1)

            # Third call, should return 2 as the random value is between 0.89 and 0.99
            self.assertEqual(random_gen.next_num(), 2)

            # Forth call, should return 3 as the random value is between 0.99 and 1.0
            self.assertEqual(random_gen.next_num(), 3)


if __name__ == '__main__':
    unittest.main()