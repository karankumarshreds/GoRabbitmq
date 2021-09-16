# Exchange type == "TOPIC"

Messages sent to a `topic` exchange MUST have a `routing-key` which is a list of words. Separated by the dots.

Example: `quick.orange.rabbit`

There can be as many words in the routing key as you like upto limit of 255 bytes. The binding key (routing key from the subscriber side) must also be in the exact same form.

Let's take an example of a key of structure: `<speed>.<color>.<species>`:

![](2021-09-17-00-31-53.png)

**NOTE** :

- `* star` can substitute for exactly one word
- `# hash` can substitute for 0 or more words

These bindings can be summarised as:

- Q1 is interested in all the orange animals.
- Q2 wants to hear everything about rabbits, and everything about lazy animals.

- A message with a routing key set to "quick.orange.rabbit" will be delivered to both queues.
- Message "lazy.orange.elephant" also will go to both of them.
- On the other hand "quick.orange.fox" will only go to the first queue, and "lazy.brown.fox" only to the second.
- "lazy.pink.rabbit" will be delivered to the second queue only once.

**Really important points:**

What happens if we break our contract and send a message with one or four words, like "orange" or "quick.orange.male.rabbit"? Well, these messages won't match any bindings and will be lost.

On the other hand "lazy.orange.male.rabbit", even though it has four words, will match the last binding and will be delivered to the second queue.
