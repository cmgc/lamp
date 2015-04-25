class Helper(object):

    def __init__(self, rows, cols, start):
        """
            start - where to start
            finish - end of path
        """
        self.limit = (rows, cols)
        self.start = start

    def horse_moves(self, point):
        orig = self.limit
        possible_places = [(1, -2), (-1, -2), (-2, -1), (-2, 1),
                           (-1, 2), (1, 2), (2, 1), (2, -1)]

        gen_moves = lambda x: (x[0] + point[0], x[1] + point[1])
        is_legal = lambda x: x[0] >= 0 and x[0] < orig[0]\
            and x[1] >= 0 and x[1] < orig[1]

        return filter(is_legal, map(gen_moves, possible_places))

    def _is_legal(self, pt, orig):
        return pt[0] >= 0 and pt[0] < orig[0] \
            and pt[1] >= 0 and pt[1] < orig[1]

    def get_short_path(self, finish, parents, result):
        if (parents[finish] != self.start):
            self.get_short_path(parents[finish], parents, result)
        result.append(finish)
