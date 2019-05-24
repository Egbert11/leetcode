# Source: https://leetcode.com/problems/construct-quad-tree/
# Author: Egbert11
# Date: 2019-05-24
#
#
# We want to use quad trees to store an N x N boolean grid. Each cell in the grid can only be true or
# false. The root node represents the whole grid. For each node, it will be subdivided into four
# children nodes until the values in the region it represents are all the same.
#
# Each node has another two boolean attributes : isLeaf and val. isLeaf is true if and only if the
# node is a leaf node. The val attribute for a leaf node contains the value of the region it represents.
#
# Your task is to use a quad tree to represent a given grid. The following example may help you
# understand the problem better:
#
# Given the 8 x 8 grid below, we want to construct the corresponding quad tree:
#
# It can be divided according to the definition above:
#
# The corresponding quad tree should be as following, where each node is represented as a (isLeaf, val) pair.
#
# For the non-leaf nodes, val can be arbitrary, so it is represented as *.
#
# Note:
#
# N is less than 1000 and guaranteened to be a power of 2.
# If you want to know more about the quad tree, you can refer to its wiki.



class Node(object):
	def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft, bottomRight):
		self.val = val
		self.isLeaf = isLeaf
		self.topLeft = topLeft
		self.topRight = topRight
		self.bottomLeft = bottomLeft
		self.bottomRight = bottomRight

	def __str__(self):
		return "val:%s, isLeaf:%s topLeft:%s topRight:%s bottomLeft:%s bottomRight:%s" % (
			self.val, self.isLeaf, self.topLeft, self.topRight, self.bottomLeft, self.bottomRight
		)

class Solution(object):
	def construct(self, grid):
		"""
		:type grid: List[List[int]]
		:rtype: Node
		"""
		return self.construct_by_range(grid, 0, len(grid), 0, len(grid[0]))

	def construct_by_range(self, grid, row_begin, row_end, col_begin, col_end):
		if row_end - row_begin <= 1:
			val = True if grid[row_begin][col_begin] == 1 else False
			node = Node(val, True, None, None, None, None)
			return node
		is_same = self.is_same(grid, row_begin, row_end, col_begin, col_end)
		if is_same:
			val = True if grid[row_begin][col_begin] == 1 else False
			node = Node(val, True, None, None, None, None)
			return node
		else:
			gap = (row_end - row_begin) / 2
			top_left = self.construct_by_range(grid, row_begin, row_end-gap, col_begin, col_end-gap)
			top_right = self.construct_by_range(grid, row_begin, row_end-gap, col_begin+gap, col_end)
			bottom_left = self.construct_by_range(grid, row_begin+gap, row_end, col_begin, col_end-gap)
			bottom_right = self.construct_by_range(grid, row_begin+gap, row_end, col_begin+gap, col_end)
			node = Node(True, False, top_left, top_right, bottom_left, bottom_right)
			return node

	def is_same(self, grid, row_begin, row_end, col_begin, col_end):
		for i in range(row_begin, row_end):
			for j in range(col_begin, col_end):
				if grid[i][j] != grid[row_begin][col_begin]:
					return False
		return True


if __name__ == "__main__":
	grid = [
		[1,1,1,1,0,0,0,0],
		[1,1,1,1,0,0,0,0],
		[1,1,1,1,1,1,1,1],
		[1,1,1,1,1,1,1,1],
		[1,1,1,1,0,0,0,0],
		[1,1,1,1,0,0,0,0],
		[1,1,1,1,0,0,0,0],
		[1,1,1,1,0,0,0,0],
	]
	root = Solution().construct(grid)
	print root