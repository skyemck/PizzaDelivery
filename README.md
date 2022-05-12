# # Pizza Delivery
## Part 1
A delivery person is delivering pizzas to an infinite two-dimensional grid of houses. She begins by delivering a pizza to the house at her starting location; a dispatcher then calls via radio and tells her where to move next. Moves are always exactly one house to the north (`^`), south (`v`), east (`>`), or west (`<`). After each move, she delivers a pizza to the house at her new location. The dispatcher has been a little spacey lately, so she sometimes ends up delivering more than one pizza to the same house.

Here are some examples:

* `>` delivers pizzas to two houses: one to the house at the starting location, and one to the house directly east of the starting location.
* `^>v<` delivers pizzas to four houses in a square; the house at the starting/ending location ends up with two pizzas.
* `^v^v^v^v^v` delivers a bunch of pizzas to some very lucky people at only two houses.

Given the string of dispatcher inputs in [this file](https://gist.github.com/mikedelorenzo-koneksa/3d273f862e919782a181c28ef5f50f54#file-pizzadeliveryinput-txt), how many houses receive at least one pizza?

## Part 2
The next day, to speed up the process, the delivery person rents a pizza-delivering goat. She and the goat begin at the same starting location, and they both deliver a pizza to this starting house. She and the goat then take turns moving based on the dispatcher's instructions.

Here are some examples:

* `^v` now delivers pizzas to three houses; The delivery person goes north and the goat goes south.
* `^>v<` now delivers pizzas to three houses; The delivery person and the goat both end up back where they started.
* `^v^v^v^v^v` now delivers pizzas to 11 houses; The delivery person treks north and the goat treks south.

Given the same string of dispatcher inputs as in [Part 1](https://gist.github.com/mikedelorenzo-koneksa/3d273f862e919782a181c28ef5f50f54#part-1), how many houses receive at least one pizza?

## Guidelines for your solution
Please provide a solution that answers [Part 1](https://gist.github.com/mikedelorenzo-koneksa/3d273f862e919782a181c28ef5f50f54#part-1) and [Part 2](https://gist.github.com/mikedelorenzo-koneksa/3d273f862e919782a181c28ef5f50f54#part-2) in a modern programming language.

By design, this is a fairly simple exercise, so we're hoping for solutions that reflect what you consider well-factored production-quality code. The best solutions will be elegant, as well as correct. Feel free to pull in any libraries you like.
