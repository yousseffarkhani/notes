First off, it all depends on how large the thing is that you are planning. Are you making a 100 line script or a 100,000 line application? Is it just you or are you working together with a team of 10 people? Is this just something for yourself or are you selling it for millions of dollars to customers worldwide? Is it just a toy or would a bug cost somebody thousands of dollars due to mistakes or wasted resources?

Regardless of which development methodology you choose (waterfall, agile, etc), you will always end up bouncing back and forth between several major tasks:

    Gathering requirements

    Designing

    Implementing

    Testing

Each later stage depends on the stage before it, but the later stages will also point out flaws in the previous stages and you will have to go back and make changes. How and when you do this is what really sets different methodologies apart, but in general, you'll be iterating over and over on those four things.

It's also important to note that this pattern applies at many different levels of abstraction. The application as a whole has requirements and needs to be designed, implemented, and tested. But thinking only at that level for a large application would be unmanageable, so we break it down into smaller parts. We think about individual problems the user wants to solve; individual tasks they need to perform; algorithms that solve each problem; individual windows and widgets for interacting with; then into the classes, methods, and variables that will make up each of those individual pieces; and finally into the line-by-line instructions that form those methods and describe those objects. At each of these levels, requirements will suggest a design that shapes an implementation that will need to be tested.

But regardless of what methodology you follow or what level of abstraction you are at, the stages are essentially the same. Let's look at each in turn.

Before you can design something, you have to understand what it is you are trying to accomplish. This is where we need requirements. These could be:

    What problem is this application supposed to solve?

    What kind of input data can I expect?

    What kind of output does it need to provide?

    What is a valid answer?

    How much time should it take?

    How many resources are available to it?

    Who are the different kinds of users who will use this?

    What purpose does each user have for using it?

    What expectations does each user have for how the software will behave?

In reality, you never know the requirements up front. You probably have a vague idea, but as you write software, test it, show it to users, and get feedback, you'll quickly realize that what you thought the requirements were, and what the user really wanted, were two different things. Even the user doesn't have a perfect idea of what they want. They'll ask for one thing, but really mean something else, or not realize that what they asked for doesn't really solve their problem. This is why we make low-fidelity prototypes, "plan to throw one away", and iterate with users. There will definitely be misunderstandings and mistakes; the faster we can get clarification and fix the problems, the less time and money we will waste getting to the final answer.

Regardless of which iteration we are on, we now have to take our best understanding of the requirements and design our application from that. This includes deciding what the modules, classes, methods, variables, algorithms, and everything else will be. This typically involves deciding "what" it will do, not necessarily "how" it will do it. For example, an alarm clock might have a description like this:

    Until the user interacts with the clock, it will display the current time

    When the user presses the "set" button, it will change into the set mode, and the time will stop counting.

    When in the set mode, each time the user presses the "minute" button, the current minute will advance by one.

    When in the set mode, pressing the "set" button again will save the current time and exit the set mode.

There's no description here of how that will be implemented, only what the observable interaction will be. At a lower level, you might then design the classes and methods:

    class Clock:

        attribute: current time (date/time)

        method: setTime(date/time)

        method: getTime(date/time)

        method: setAlarm(date/time)

        method: connectAlarm(callback)

And you would keep on going down and down the layers of abstraction until it you reach the point where it's obvious how to write it and it's just a matter of typing it in.

That brings us to the implementation. Now that you have a solid plan for what the code will be, it's time to write it.

As you write the code, you will also need to test that it does what the requirements said it should do. At the lowest level, writing unit tests is a common way for automating the testing of the code. You can also write automated tests for entire classes, as well as putting together multiple classes/modules and testing that they work properly together (integration testing), and even testing the entire flow of the application as a user would use it for a specific task (system-level testing). At the higher levels it might be difficult to write code to full automate this, so the developer and/or a tester might step through tests manually, perhaps using test harnesses or other supporting structures to help them test it more consistently and efficiently. The eventual, true test, is giving it to users.

At each step along the way, you'll find mistakes. Unit tests will point out a missing case or a null pointer reference. Integration tests will point out mismatches between data formats or calling patterns. While trying to make the design, you may come up with questions that aren't answered by the requirements, forcing you to go back and elaborate on the requirements, perhaps by making a prototype and showing it to a user to get feedback. Sometimes during implementation you'll realize that there is a contradiction in the design and it can't be implemented as designed, so the design needs to be changed. Or you'll realize that the design doesn't actually fulfill the requirements. Whatever the reason, you will iterate many times, jumping back and forth between stages and levels of abstraction. Or maybe you'll give it to the user and they'll say, "that's not what I wanted" and you'll have to go all the way back to changing the requirements and the changes will ripple down throughout the design and implementation.

I'm not sure I can actually say how much time I spend on "conception", but a rough estimate for me would be an even 3-way split between: requirements and design, implementation, testing and debugging. Note that when people make estimates for a project, they almost always think only about the implementation and they completely overlook the rest, but it makes up two-thirds of the time spent working on something.

For techniques, I start something like talking to a person on the phone or in person, preferably with a whiteboard, and we talk and sketch out our ideas. This is great for both requirements and design. Depending on how many questions are left unanswered, I might then make a prototype and try it out. For a UI, this might be as simple as drawing windows on a piece of paper and asking the user to tell me what they want to do and me explaining what happens. Or I might put together a very rough program with no error handling or pleasantries that is the bare minimum needed to test my assumption or demonstrate my idea. I'll usually make notes in a notebook about what we came up with, though I would rarely make a formal document stating requirements or the design (though it entirely depends on how formal the customer wants to be).

Finally, does it make a difference if it's for work or personal? Absolutely. My personal projects are typically small, non-essential, and I am my own user, so I (mostly) know my requirements upfront (though even a designer/developer doesn't always consciously know what they want). I'll be a lot less worried about strange corner cases, or integration with 3rd party tools, or ensuring that the user can't accidentally destroy all their hard work. Basically, I'll be sloppy on a personal project. But for work, you have to consider all these things.

Also, the iteration times are much smaller for a personal project. If I find a bug in my own software, I'll dig into it right away and likely have it fixed in less than an hour. If a customer finds a defect in my software, I might not hear about it for weeks, and it might be days for a single iteration with them. Because of this, I'll be a lot more formal and careful when iterating with a user (since a mistake could cost days), whereas I'll be far less careful when iterating myself (since a mistake only costs minutes). This is exactly the same thing as how finding bugs in development is far cheaper than finding bugs in a released product.

Basically, developing personal software is in many ways different than developing for work or for somebody else in general. The bigger the project, the more disconnected I am from the end user, and the more important it is that things work, the more I will spend on designing and testing.

That said, once you get used to developing things in that careful and formal manner, you find yourself doing it for your own projects, it makes things much faster and smoother. The whole point of that extra work is that it saves you far more time later. Perhaps the costs aren't as high in a personal project, but planning, designing, and testing are still valuable investments in personal projects.

I hope that answered your questions.

If you want to know more, the book Code Complete has a great discussion on this.