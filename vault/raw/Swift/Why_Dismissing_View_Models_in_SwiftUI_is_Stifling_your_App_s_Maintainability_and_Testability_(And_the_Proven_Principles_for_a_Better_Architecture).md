---
title: "Why Dismissing View Models in SwiftUI is Stifling your App’s Maintainability and Testability (And the Proven Principles for a Better Architecture)"
source: "https://matteomanferdini.com/swiftui-viewmodel/?utm_source=x&utm_medium=social&utm_campaign=swiftui-viewmodel&utm_id=content-calendar-2026-04-20"
author:
  - "[[Matteo Manferdini]]"
published: 2025-09-02
created: 2026-04-25
description: "If you’ve been working with SwiftUI, you’ve likely noticed that your views start pretty simple but then balloon into large, unmaintainable monoliths that are hard to preview and test. While there are several techniques to keep SwiftUI views modular and reusable, some problems are architectural in nature and can only be addressed by following proven ... Read more"
tags:
  - "clippings"
---
If you’ve been working with SwiftUI, you’ve likely noticed that your views start pretty simple but then balloon into large, unmaintainable monoliths that are hard to preview and test.

While there are several techniques to [keep SwiftUI views modular and reusable](https://matteomanferdini.com/swiftui-views/), some problems are architectural in nature and can only be addressed by following proven software design principles.

Particularly, view models are an essential component to guarantee testability, maintainability, and code reuse across views.

![](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

## FREE GUIDE - SwiftUI App Architecture: Design Patterns and Best Practices

MVC, MVVM, and MV in SwiftUI, plus practical lessons on code encapsulation and data flow.

## Table of contents

###### Chapter 1

## The Software Design Principles behind MVVM

**In this chapter:**

## Monolithic SwiftUI views are hard to maintain and unit test

The core issue of monolithic types is that existing code invites the accumulation of related code, which leads to lower cohesion.

In SwiftUI, that happens when **views start accumulating** ***app business logic***, often in the form of:

- networking code,
- error handling, and
- state transition logic.

These, in turn, make the code of SwiftUI views tightly coupled, making them hard to maintain, difficult to preview, and impossible to unit test.

The solution to these problems is to adopt an architectural design pattern, such as MVVM, that distributes responsibilities across well-defined layers.

## Should you use view models in SwiftUI, or is the MVVM pattern outdated?

MVVM has been used by numerous developers for over two decades on various platforms, including iOS. However, you often find online arguments against using MVVM in SwiftUI, such as:

- It’s [the current year](https://en.wikipedia.org/wiki/Chronological_snobbery), and you should not use [outdated design patterns](https://en.wikipedia.org/wiki/Appeal_to_novelty).
- View models allegedly make your code harder and offer no benefits.
- View models allegedly conflict with Apple’s frameworks.
- [Our Apple overlords](https://en.wikipedia.org/wiki/Argument_from_authority) allegedly don’t want you to use MVVM.

This leaves many developers wondering whether they should use view models in their apps or abandon them completely for a new, trendy pattern like the MV pattern.

I have already explained in detail [how to use MVVM in SwiftUI](https://matteomanferdini.com/swiftui-mvvm/), and you can delve deeper into the topic using [my free guide on app architecture.](https://matteomanferdini.com/swiftui-app-architecture/) I’ll use this article to address those criticisms.

###### Note

My arguments rest on several software design principles, including:

- [Encapsulation](https://en.wikipedia.org/wiki/Encapsulation_\(computer_programming\));
- The [Law of Demeter](https://en.wikipedia.org/wiki/Law_of_Demeter);
- [Separation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns);
- [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself);
- [SOLID](https://en.wikipedia.org/wiki/SOLID);
- [cohesion](https://en.wikipedia.org/wiki/Cohesion_\(computer_science\)) and [coupling](https://en.wikipedia.org/wiki/Coupling_\(computer_programming\));
- The twenty-three [design patterns](https://en.wikipedia.org/wiki/Design_Patterns) documented by the [Gang of Four](https://ieeexplore.ieee.org/document/8009931).

My arguments require believing that **these are timeless principles** that still apply, regardless of the programming language or UI framework used, as well as understanding the problems these principles aim to solve.

This is definitely not an opinion shared by everyone, as exemplified by this comment I got on LinkedIn:

[![](data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%201042%20342'%3E%3C/svg%3E)](https://matteomanferdini.com/wp-content/uploads/2025/09/a-linkedin-comment-against-the-solid-principles.png)

You might also have stumbled upon [this long rant](https://developer.apple.com/forums/thread/699003) against MVVM on Apple’s developer forums, where the author calls the SOLID principles and unit testing “a cancer on iOS development”.

If you share those opinions, you probably won’t get much value out of this article, so I’ll hopefully save you time by putting this note up front.

###### Chapter 2

## View State Transitions and Imperative Logic

**In this chapter:**

![](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

## FREE GUIDE - SwiftUI App Architecture: Design Patterns and Best Practices

MVC, MVVM, and MV in SwiftUI, plus practical lessons on code encapsulation and data flow.

## The state transitions of a view are driven by imperative code, despite SwiftUI’s declarative nature

To avoid being accused of criticizing something nobody said, we will start with an example I took from an online article opposing view models.

This is typically the code for a view that fetches its data from a REST API, a pattern I have seen in many code samples.

```swift
struct ItemsView: View {
    @State private var viewState: ViewState = .loading
    
    enum ViewState {
        case loading
        case loaded([Item])
        case error(String)
    }
    
    var body: some View {
        NavigationStack {
            List {
                switch viewState {
                    case .loading: // Display the loading view
                    case .loaded(let array): // Display the list of items
                    case .error(let string): // Display the error view
                }
            }
            .task { await load() }
        }
    }
}
```

We don’t know yet how the view will transition from `.loading` to `.loaded(_:)` or `.error(_:)`. However, we can already be certain that some code will express that logic.

Thus, **the view’s declarative code will be driven by imperative logic** that changes its state.

###### Misconception

Tangentially, the `ViewState` enumeration in the code above is a bad way to represent the states of a view.

Some developers seem to believe it is an appropriate implementation of the [state pattern](https://en.wikipedia.org/wiki/State_pattern), but it clearly isn’t, as the state pattern requires subclassing.

Generally, enumerations are **a poor abstraction mechanism** **to represent the entire state** of an entity.

Enumerations obviously have their application in representing mutually excluding options in individual stored properties. However, since they are [sum types](https://en.wikipedia.org/wiki/Tagged_union), they can’t represent overlapping states as efficiently as [product types](https://en.wikipedia.org/wiki/Product_type).

As such, **enumerations should not encapsulate the entire state of a view**. For example, the code above cannot handle a case where the view still needs to display the previous list of items when a reloading error occurs.

Moreover, using enumerations to abstract the internal state of an entity and its transitions violates the [Open-Closed principle](https://en.wikipedia.org/wiki/Open%E2%80%93closed_principle).

This becomes evident in the example above when you try to implement the overlap I just mentioned, for example, by adding an `Item` array to the `error` case.

To be fair, in such a simple piece of code, you can get away with pretty much anything. **Any violation of a software development principle is unlikely to be a significant issue when your views are small enough**.

However, those principles exist to prevent problems as code becomes increasingly complex, and are not invalidated by such simple examples.

Therefore, while the examples in this article will be deliberately simple to present my arguments in the most straightforward way possible, I trust your experience with building real apps to help you imagine how each problem would worsen in more complex scenarios.

## The imperative logic of a SwiftUI view is untestable because of the framework’s internal implementation

The loading of the view’s content generally happens in one or more private methods. This is where we find the imperative code that updates the view’s state.

```swift
private extension ItemsView {
    func load() async {
        do {
            let (data, _) = try await URLSession.shared
                .data(from: URL(string: "example.com")!)
            let items = try JSONDecoder().decode([Item].self, from: data)
            viewState = .loaded(items)
        } catch {
            viewState = .error("Error message")
        }
    }
}
```

Unfortunately, **this function is untestable** because of how SwiftUI works and how the `@State` property wrapper is implemented.

Even though the `ItemsView` is a Swift structure, creating an individual view value in a unit test and then inspecting the content changes of the `viewState` property does not work.

Detractors of MVVM think that this does not matter and recommend moving the networking code into a shared object for testability.

```swift
@Observable class NetworkController {
    func fetchItems() async throws -> [Item] {
        let (data, _) = try await URLSession.shared
            .data(from: URL(string: "example.com")!)
        return try JSONDecoder().decode([Item].self, from: data)
    }
}

struct ItemsView: View {
    @State private var viewState: ViewState = .loading
    @Environment(NetworkController.self) private var networkController

    // ...
}

private extension ItemsView {
    func load() async {
        do {
            let items = try networkController.fetchItems()
            viewState = .loaded(items)
        } catch {
            viewState = .error("Error message")
        }
    }
}
```

###### Note

Such a shared object is called a *controller* in MVC, but is also sometimes referred to as a *service*, a *manager,* a *store*, or a *client*.

However, the name doesn’t change the nature of the object, which is to be shared across multiple SwiftUI views, either as a [singleton](https://matteomanferdini.com/swift-singleton/) or through the SwiftUI environment.

The `fetchItems()` function is now testable. However, architecturally, this hardly solves anything.

Pushing code like this into a shared object typically violates the *DRY (Don’t Repeat Yourself) principle*, as the different networking methods end up having code that is structurally similar.

The common solutions to such repetition you find online usually violate:

- The *Open-Closed principle*, since code is often riddled with conditional statements that must be updated for every request.
- The *Dependency Inversion principle*, as trying to remove repetition through excessive parameterization leads to bottom-up design, with the implementation details influencing the clients of the interface.
- The *Interface Segregation principle,* since the monolithic object exposes several methods to clients, with the risk of creating indirect coupling (however, this is easily solved using protocols).

I illustrate all these problems and the proper way to address them in my article on making [REST API Calls in Swift](https://matteomanferdini.com/swift-rest-api/#protocol-oriented-networking-layer-architecture).

## A view’s state transitions are a crucial part of the app’s business logic

Moving networking logic into a shared object **does not improve the testability of the imperative code that drives the state changes of a view**. In our example, the state transition code of the `ItemsView` remains within its `load()` method.

Conditional flow statements like `if`, `switch`, and `do` are an expression of the *app’s business logic*. However, in this case, such logic does not belong to shared objects, since it is particular to this view. Namely, the view’s logic defines:

- What user interface is displayed according to the view’s state.
- Which API endpoints are called to load the view’s data, and in which order if there is more than one.
- How the view state changes with time, i.e., how it moves from `.loading` to `.loaded(_:)` or `error(_:)`.
- What state transitions should not be possible, e.g., going from `error(_:)` to `.loaded(_:)`.

**There is nothing intrinsically wrong with having this code inside a view.** However, as we have seen, you can’t [unit test](https://matteomanferdini.com/ios-unit-testing/) that logic, which is particularly important, especially in views with logic more complex than such a toy example.

Moreover, such a view is also **difficult or impossible to preview in Xcode,** as it usually requires an active internet connection and the instantiation of a database or the entire app networking stack.

### But I ain’t testing my views

Cool. If your argument is that you don’t care about testing the code of your views, then you might not need a view model.

Even [Wikipedia acknowledges](https://en.wikipedia.org/wiki/Software_design_pattern) that (emphasis mine):

A software design pattern is a general, reusable solution to a **commonly occurring problem**.

If you don’t have a problem, you don’t need its solution. It is a mistake to blindly adhere to any specific design pattern when it’s not required.

However, **“I don’t care” is not an argument against MVVM**. It’s not even an argument against others caring.

###### Note

**I don’t write view models for sports**, and I have plenty of views in my apps with logic that is simple enough not to require unit testing. I am also the first to argue that excessive unit testing can become a liability.

Moreover, [most views do not need a view model](https://matteomanferdini.com/swiftui-massive-reusable-views/) because they are at the wrong architectural level.

### The Apple development community has historically been resistant to unit testing

At my first job, we used C# on Windows and attended conference talks on unit testing. Meanwhile, in my spare time, I was developing for the Mac, and I read blogs from prominent Mac developers [who argued against unit testing](http://blog.wilshipley.com/2005/09/unit-testing-is-teh-suck-urr.html).

Support for unit testing was added to Xcode 5 only in 2013. That might sound like a long time ago now, but I started developing apps for the Mac in 2005, so I worked in an environment without unit testing for 8 years.

I have worked at startups where nothing was tested. However, I also worked at large companies where the testing requirements were stringent.

Moreover, I worked at companies that wanted to introduce unit testing into codebases that made it impossible, and the only solution was a complete rewrite, which management never allowed to happen. So we had to suck it up and live with bugs that were unfixable.

Pick your poison.

###### Note

Unit testing is not the only argument in favor of view models. There are [specific cases and patterns](https://matteomanferdini.com/swift-download-file-from-url/) that require the implicit identity of objects, which SwiftUI views, being structures, cannot provide.

I will also provide a more complex example in the last section of this article.

## Aggregate models should not contain view-specific logic, as they are shared across several views

Some developers who argue against view models acknowledge the need to embed some of the app’s business logic into classes that should typically reside in a view model for testability.

These are sometimes called *aggregate models,* and you can find examples in Apple’s sample apps [Fruta](https://developer.apple.com/documentation/appclip/fruta_building_a_feature-rich_app_with_swiftui) and [Food Truck](https://developer.apple.com/documentation/swiftui/food_truck_building_a_swiftui_multiplatform_app/).

**Aggregate models are also shared objects**, sometimes sitting just above controllers (or whatever you call those). They store data and encapsulate logic shared across several views.

Like other shared objects, they are distributed through the SwiftUI environment (although, curiously, in the Food Truck app, they are passed through initializers in the entire view hierarchy).

```swift
class Model: ObservableObject {
    // ...

    @Published var favoriteSmoothieIDs = Set()

    // ...
}

struct FrutaApp: App {
    @StateObject private var model = Model()

    var body: some Scene {
        WindowGroup {
            ContentView()
                .environmentObject(model)
        }
        // ...
    }
}

struct FavoriteSmoothies: View {
    @EnvironmentObject private var model: Model

    var favoriteSmoothies: [Smoothie] {
        model.favoriteSmoothieIDs.map { Smoothie(for: $0)! }
    }

    var body: some View {
        SmoothieList(smoothies: favoriteSmoothies)
        // ...
    }
}
```

And that’s the crux of the matter. Being, like controllers, shared objects that encapsulate shared logic, **they should not contain view-specific code**, as I previously discussed.

###### Chapter 3

## Implementing and Testing View Models

**In this chapter:**

![](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

## FREE GUIDE - SwiftUI App Architecture: Design Patterns and Best Practices

MVC, MVVM, and MV in SwiftUI, plus practical lessons on code encapsulation and data flow.

## Implementing view models as observable objects

When implementing a view model, the general approach is to use an `@Observable` class on the main actor, as all data changes that trigger a UI update must occur on the main thread.

```swift
@MainActor @Observable class ViewModel {
    var items: [Item] = []
    var isLoading = false
    var errorOccurred = false

    private let networkController: NetworkController

    init(networkController: NetworkController) {
        self.networkController = networkController
    }

    func load() async {
        isLoading = true
        defer { isLoading = false }
        do {
            items = try await networkController.fetchItems()
        } catch {
            errorOccurred = true
        }
    }
}
```

I removed the `ViewState` enumeration, as it was a bad practice, as I mentioned in a note above.

###### Note

With the introduction of [approachable concurrency](https://github.com/swiftlang/swift-evolution/blob/main/visions/approachable-concurrency.md) in Swift 6.2, the `@MainActor` attribute is not needed anymore, as all code runs on the main actor by default.

This makes the view logic finally testable. We can stub the `NetworkController` to remove the need for a network connection.

```swift
class NetworkControllerStub: NetworkController {
    override func fetchItems() async throws -> [Item] {
        return [ Item() ]
    }
}

@MainActor @Test func example() async throws {
    let viewModel = ViewModel(networkController: NetworkControllerStub())
    await viewModel.load()
    #expect(!viewModel.isLoading)
    #expect(!viewModel.errorOccurred)
    #expect(viewModel.items.count == 1)
}
```

###### Note

In specific cases, view models could also be implemented as Swift structures. However, there are many scenarios that require an object’s built-in identity, for example:

- Delegation.
- Escaping closures capturing a mutating `self` parameter, like the trailing closure of a `Task`.
- Design patterns that require reference types, like the state pattern.
- Generic views with view models that can be extended through subclassing. (You can find an example of this in the last section of this article.)
- To easily create test doubles through method overrides.

## You still need unit tests even when you use UI tests

Another argument against view models is that you can always test your views’ code using Xcode’s UI tests.

That’s not necessarily wrong. However, unit testing and UI testing are different and complementary, and the latter does not replace the former.

UI tests sit at the top of the [testing pyramid](https://martinfowler.com/articles/practical-test-pyramid.html), as illustrated in [Apple’s documentation on testing](https://developer.apple.com/documentation/xcode/testing).

[![](data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%202560%201609'%3E%3C/svg%3E)](https://matteomanferdini.com/wp-content/uploads/2025/09/testing-pyramid-scaled.jpeg)

**UI tests test an entire user flow**. While they are a useful tool, they involve a significant portion of your codebase, are more difficult to set up, slower to run, less precise, and prone to breaking. That’s why they are at the top of the pyramid, above integration tests.

Unit tests, on the other hand, are easier to write and maintain, quicker to run, and allow you to **test granular behavior and edge cases** that are harder or impossible to test through a UI test.

For example, I would never use a UI test to replace the simple unit test I showed above.

For these reasons, you should have many more unit tests than UI tests in your codebase.

## View models should be instantiated inside a view

Some of the arguments surrounding view models center on their instantiation.

[According to Apple’s documentation](https://developer.apple.com/documentation/swiftui/state#Store-observable-objects), `@Observable` objects contained by `@State` properties should be created using the [task(priority:\_:)](https://developer.apple.com/documentation/swiftui/view/task\(priority:_:\)) modifier, which is triggered only when the view appears on screen.

```swift
struct ItemsView: View {
    @State private var viewModel: ViewModel!
    @Environment(NetworkController.self) private var networkController

    var body: some View {
        NavigationStack {
            Group {
                if let viewModel {
                    if viewModel.isLoading {
                        // Display the loading view
                    } else if viewModel.errorOccurred {
                        // Display the error view
                    } else {
                        List(viewModel.items) { item in
                            // Display the list of restaurants
                        }
                    }
                }
            }
            .task {
                guard viewModel != nil else { return }
                viewModel = ViewModel(networkController: networkController)
                await viewModel.load()
            }
        }
    }
}
```

###### Note

I would not normally place a `NavigationStack` inside this view. But that’s a different topic altogether for another article. I added one here just for the sake of the example.

Some developers argue against this pattern because it requires making the view model property optional, with all the annoying related unwrapping and checks for `nil`. These can only be partially relieved by using an implicitly unwrapped optional (which is not necessary if you prefer a normal optional).

I agree that the extra code introduced by optionals can be annoying. However, other solutions are worse, as I will show below.

###### Note

The only architecturally sound alternative is to [instantiate view models inside coordinators](https://matteomanferdini.com/mvvm-coordinator-swiftui/) and then inject them into their respective view through an initializer.

However, I would recommend that pattern only in complex SwiftUI apps. Coordinators can be useful for complex navigation and testing, but I wouldn’t introduce them only for the sake of removing Swift optionals from views.

### Creating a view model in the view initializer works with the @StateObject property wrapper, but not with @State

One option is to instantiate a view model in the view’s initializer. However, this works when using the old approach to observable objects, which requires

- a class conforming to `ObservableObject` with `@Published` properties;
- a `@StateObject` stored property in the view.

The reason this works is that [the initializer](https://developer.apple.com/documentation/swiftui/stateobject/init\(wrappedvalue:\)) of the `@StateObject` property wrapper has an escaping autoclosure, which SwiftUI runs only once.

```swift
@MainActor class ViewModel: ObservableObject {
    // ...
}

struct ItemsView: View {
    @StateObject private var viewModel: ViewModel

    init() {
        // This view model is instantiated only once, regardless of how
        // many times the view's initializer is called.
        self._viewModel = StateObject(wrappedValue: ViewModel())
    }

    // ...
}
```

However, the `@State` property wrapper currently does not have such an initializer. Following the same approach **resets the view model every time the view initializer runs**.

```swift
@MainActor @Observable class ViewModel {
    // ...
}

struct ItemsView: View {
    @State private var viewModel: ViewModel

    init() {
        // This view model is recreated every time the view's initializer
        // is executed because of a user interface update.
        self._viewModel = State(wrappedValue: ViewModel())
    }

    // ...
}
```

### Passing a view model or other dependencies to a view initializer violates several design principles

Even when using the old `@StateObject` property wrapper, instantiating a view model in the initializer of a view is only acceptable when the view model does not need to access any environment objects.

**Environment objects are available only after a view initializer runs**, when the `body` of the view is executed. That is why Apple documentation recommends using the `task(priority:_:)` modifier.

Some developers prefer to set up all required properties at initialization to avoid using optionals. While I share the sentiment, this leads to other problems.

To set up a view model with dependencies using an initializer, the dependencies or the view model instance itself must be passed as parameters. However, this pattern again violates several principles.

- If the view model is passed as a parameter, it violates the *Single Responsibility principle* as the preceding view gets coupled to a view model it does not own or need.
- If a shared object is passed as a parameter, it defeats the purpose of the environment, as the preceding view must access an environment object that it might not need otherwise (violating the *Single Responsibility principle again*).
- It violates the *DRY principle* if the view requiring the view model needs to be created in more than one place in the codebase.
- If the preceding view created several views that require a view model, this is likely to violate the *Open-Closed principle*.
- When using navigation views, the view model is created too early, even for views that are never reached. Besides allocating useless objects, this can also potentially trigger unnecessary database access or network requests.
- In a `List` with navigation links with a view with a view model as a `destination`, multiple instances would be created unnecessarily. (These last two points were a problem in the first days of SwiftUI, before the `@StateObject` property wrapper was introduced).

###### Note

The last problem can be avoided by using navigation links [initialized with values](https://developer.apple.com/documentation/swiftui/navigationlink/init\(_:value:\)) managed by the [navigationDestination(for:destination:)](https://developer.apple.com/documentation/swiftui/view/navigationdestination\(for:destination:\)) view modifier, which is located higher in the view hierarchy.

###### Chapter 4

## How MVVM Interacts with Apple Frameworks

**In this chapter:**

![](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

## FREE GUIDE - SwiftUI App Architecture: Design Patterns and Best Practices

MVC, MVVM, and MV in SwiftUI, plus practical lessons on code encapsulation and data flow.

## Is Apple actively against MVVM, or do they endorse it?

On the web, you can find many claims that MVVM goes against Apple frameworks, that Apple does not recommend MVVM, or even that they are explicitly against it.

I’ll start by saying that I don’t find that line of argument particularly interesting, as it’s just an [argument from authority](https://en.wikipedia.org/wiki/Argument_from_authority). **Any argument for or against MVVM should rest solely on its merits**. Apple is not the final arbiter of your app’s architecture.

However, since some do pursue this line of argument, I’ll address those claims as well.

The provided “proof” is usually that in some WWDC videos about SwiftUI, Apple “barely” mentions view models. For starters, barely is not zero, which would be the number if Apple were actively against view models.

**We can easily disprove those claims by searching for “viewmodel”** in the Apple Developer app and seeing it appear in several sessions, many as late as 2024 and 2025 (SwiftUI was introduced in 2019).

[![](data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%202560%201605'%3E%3C/svg%3E)](https://matteomanferdini.com/wp-content/uploads/2025/09/the-use-of-view-models-in-wwdc-videos-scaled.jpeg)

In fact, I would argue, if you care for these arguments anyway, that Apple introduced the `@StateObject` property wrapper [at WWDC20](https://developer.apple.com/videos/play/wwdc2020/10041) to explicitly allow developers to implement view models if they wanted to, since many were complaining about it (but that’s merely my opinion).

While **this is not an endorsement of MVVM by Apple**, it addresses the claims that Apple is somehow opposed to the use of view models or that they have no place in SwiftUI apps.

## Is MVVM incompatible with SwiftData and other Apple frameworks?

Another common argument against view models, particularly in relation to Apple frameworks, is that [MVVM does not fit SwiftData](https://matteomanferdini.com/swiftdata-mvvm/).

Particularly, to my knowledge, the argument boils down to the fact that **the `@Query` macro can only be used inside views**, which is an argument I don’t find very compelling.

That argument rests solely on the idea that, for MVVM to be implemented correctly, any `@Query` property should be moved into a view model, along with the rest of the view logic.

That, however, is not necessary. For example, we can update our view model example by adding the template SwiftData code that comes with a new Xcode project.

```swift
@MainActor @Observable final class ViewModel {
    // ...
    private let modelContext: ModelContext

    init(modelContext: ModelContext, networkController: NetworkController) {
        self.modelContext = modelContext
        self.networkController = networkController
    }

    func load() async {
        isLoading = true
        defer { isLoading = false }
        do {
            // The view model can update the model context while the @Query
            // stored property remains in the view.
            let items = try await networkController.fetchItems()
            try modelContext.transaction {
                for item in items {
                    modelContext.insert(item)
                }
            }
        } catch {
            errorOccurred = true
        }
    }
}
```

**This view model does not require a `@Query` property**. You can leave that in the view, driving the interface updates when the underlying context changes.

```swift
struct ItemsView: View {
    @State private var viewModel: ViewModel!
    @Query private var items: [Item]
    @Environment(\.modelContext) private var modelContext
    @Environment(NetworkController.self) private var networkController

    var body: some View {
        NavigationStack {
            Group {
                if let viewModel {
                    if viewModel.isLoading {
                        // Display the loading view
                    } else if viewModel.errorOccurred {
                        // Display the error view
                    } else {
                        List(items) { item in
                            // Display the list of restaurants
                        }
                    }
                }
            }
            .task {
                guard viewModel != nil else { return }
                viewModel = ViewModel(
                    modelContext: modelContext,
                    networkController: networkController
                )
                await viewModel.load()
            }
        }
    }
}
```

### The SwiftData model context is the app’s single source of truth

If I had to guess, I think some developers might feel they need to move every `@Query` property into a view model to adhere to the *single source of truth* principle that Apple stressed since the release of SwiftUI.

However, that principle is not broken when leaving `@Query` properties in the view. **The single source of truth is the underlying SwiftData `ModelContext`**. A `@Query` property is a mere convenience to retrieve data from such a context.

In the view model example above, accessing the array of items is not needed. However, if you need that, you can do so with a computed property using a `FetchDescriptor`.

```swift
@MainActor @Observable final class ViewModel {
    // ...

    // This is a read-only computed property.
    private var items: [Item] {
        (try? modelContext.fetch(FetchDescriptor())) ?? []
    }

    //...

    // Only the model context is modified. There is no need to update 
    // a stored property since the view still has a @Query property.
    func deleteItems(offsets: IndexSet) {
        for index in offsets {
            modelContext.delete(items[index])
        }
    }
}
```

That is not a crime, like someone might want you to think.

Moreover, such a property is only for the view model code. That’s why I made it `private`. **Nothing forces you to access it from the view** once it has been created. You can keep using the `@Query` macro.

###### Chapter 5

## The MVC and MVVM Design Patterns in SwiftUI

**In this chapter:**

![](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

## FREE GUIDE - SwiftUI App Architecture: Design Patterns and Best Practices

MVC, MVVM, and MV in SwiftUI, plus practical lessons on code encapsulation and data flow.

## Apple used to explicitly endorse a variation of MVC equivalent to MVVM, which was baked into UIKit

Does Apple ultimately recommend, or at least silently endorse, any specific design patterns?

**Nowadays, Apple engineers don’t explicitly recommend any architectural pattern**, since it’s not their job. I searched far and wide, and I could not find any explicit endorsement.

The only thing I can recall is a third-hand, recycled rumor that someone claimed to have been told by an Apple engineer in person that MVVM naturally fits SwiftUI, but I couldn’t even find the source of that rumor, and it would be a weak argument anyway.

###### Note

I found an article about MVVM on LinkedIn written by an engineer currently working at Apple. However, the article was published before he joined Apple, so, in my opinion, it does not hold much weight.

Apple provides flexible frameworks, **leaving architectural decisions to developers** who know best what fits their requirements.

### The SwiftUI data flow diagram is nothing else than MVC

Critics of MVVM take this silence to claim that Apple’s WWDC talks silently endorse their personal interpretation of SwiftUI architecture, which they sometimes refer to as the [*MV pattern*](https://matteomanferdini.com/swiftui-mv-pattern/).

I think that claim is misleading. However, if we want to play that game, **I can present stronger evidence of Apple endorsing the MVC pattern** rather than anything else.

An often-cited source is a chart from Apple’s WWDC19 talk, [“Data Flow Through SwiftUI](https://developer.apple.com/videos/play/wwdc2019/226),” where everyone marvels at the utter simplicity of this “new” pattern when compared to others.

The irony is that this is the same chart [on page 10 of this document, which introduces MVC](https://folk.universitetetioslo.no/trygver/2003/javazone-jaoo/MVC_pattern.pdf), and was created 40 years before the introduction of SwiftUI by [Trygve Reenskaug](https://en.wikipedia.org/wiki/Trygve_Reenskaug).

[![](data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%201440%20900'%3E%3C/svg%3E)](https://matteomanferdini.com/wp-content/uploads/2025/09/data-flow-in-swiftui-and-mvc.jpeg)

###### Note

The images in this section are slides from a video in my advanced course, *Scalable Networking Architecture*, titled “One Pattern to Rule Them All: Why MVC, MVVM, MV, and other variations are the same pattern.”

However, you don’t even need to dig that far, as the same diagram can currently be found, upside-down, on [Wikipedia’s page for the MVC pattern](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller).

[![](data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%202560%201605'%3E%3C/svg%3E)](https://matteomanferdini.com/wp-content/uploads/2025/09/the-mvc-pattern-diagram-on-wikipedia-scaled.jpeg)

If you want to claim that Apple endorses a design pattern for SwiftUI, that’s MVC.

This is not a surprise since [Apple has explicitly endorsed the MVC pattern in the past](https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Model-View-Controller/Model-View-Controller.html#//apple_ref/doc/uid/TP40010810-CH14-SW1). In fact, **MVC is such a fundamental pattern in software development** that any other successor is simply a rehash of the same ideas.

### The dawning of MVVM in Apple documentation and iOS

If you want to go a step further, you can even see an equivalent to the MVVM pattern appear in Apple’s documentation and frameworks.

If you read Apple’s guide about MVC carefully, you will find a section explaining how the controller layer can be split into *model controllers,* which are primarily concerned with the model layer, and *view controllers,* which are primarily concerned with the view layer.

The concept of view controllers was introduced long before the iPhone existed and was later explicitly integrated into UIKit. It does not take an expert in graph theory to recognize that **the two patterns are equivalent**.

[![](data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%201440%20900'%3E%3C/svg%3E)](https://matteomanferdini.com/wp-content/uploads/2025/09/the-mvvm-pattern-appearing-in-apple-documentation.jpeg)

Does this translate to SwiftUI? Some claim it doesn’t, and it’s fair to say that, since UIKit and SwiftUI are so different, the former does not affect the latter.

However, I believe that **architectural lessons do not disappear when switching to a new UI framework**, and I provided ample evidence in this article. In the end, you decide what’s best for your apps.

## The MV pattern is an anti-pattern because it’s too prescriptive

MVC and MVVM are examples of a [multitier architecture](https://en.wikipedia.org/wiki/Multitier_architecture) that has typically three or four layers. **Such an architecture can be expanded as needed**, going, for example, from the three layers of MVC to the four of MVVM. In fact, I propose adding a fifth layer, [which I refer to as the root layer](https://matteomanferdini.com/swiftui-massive-reusable-views/).

[![](data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%201440%20900'%3E%3C/svg%3E)](https://matteomanferdini.com/wp-content/uploads/2025/09/extending-the-mvc-and-mvvm-patterns.jpeg)

As I showed above, the diagram Apple provided for SwiftUI’s data flow is based on the MVC pattern. From this perspective, one could argue that the MV pattern, which also uses controller-like shared objects, is essentially the same as MVC, albeit with a different name.

However, I believe their nature is different, and MV is an [anti-pattern](https://en.wikipedia.org/wiki/Anti-pattern). The reason is that it does not merely propose a starting point, like MVC. Instead, it was conceived as a compression of layers, rather than an expansion, explicitly in opposition to MVVM.

[![](data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%201440%20900'%3E%3C/svg%3E)](https://matteomanferdini.com/wp-content/uploads/2025/09/all-popular-architectural-design-patterns-are-just-mvc.jpeg)

As I have argued throughout the article, **forcefully compressing responsibilities into the view layer leads to a violation of several fundamental software design principles**.

While view models are not always required, and you might end up with an MVC architecture in some parts of your app, the MV pattern restricts you to this approach and recommends avoiding view models even when they are useful.

## A design pattern for real-world SwiftUI applications

While this was a lengthy article with numerous examples, I kept them all simple to present each argument individually. However, the nuances of real-world software development are often lost when using toy examples.

Therefore, I would like to conclude the article with a more complex example taken from the final modules of my *Scalable Networking Architecture* course. This will demonstrate how the MVVM pattern is not only necessary to adhere to advanced software design principles, but also to implement standard code reuse practices such as [modularity](https://en.wikipedia.org/wiki/Modular_programming), [reusability](https://en.wikipedia.org/wiki/Reusability), and [composition](https://en.wikipedia.org/wiki/Composition_over_inheritance).

In the course, I built a networked app that interacts with the [GitHub REST API](https://docs.github.com/en/rest?apiVersion=2022-11-28), handling authentication, concurrent requests, rate limiting, offline caching, and error handling.

The app displays screens with a list of entities retrieved from the API. While I limit the example to users and repositories, the API allows fetching all sorts of lists, e.g., organizations, teams, issues, or pull requests.

[![](data:image/svg+xml,%3Csvg%20xmlns='http://www.w3.org/2000/svg'%20viewBox='0%200%202560%201648'%3E%3C/svg%3E)](https://matteomanferdini.com/wp-content/uploads/2025/09/a-list-view-displaying-rows-for-users-and-repos-scaled.jpeg)

While they differ in their appearance, these list screens obviously have a lot in common, such as loading and refreshing their content, and paging.

While different entities have different endpoint URLs, the sequence of API calls to fetch the list content is similar, requiring:

1. A network request to fetch a page in the list.
2. A series of parallel requests to fetch the toggle status of each element, which is provided by a different endpoint, i.e., if a user is followed or if a repository is starred.

However, each list also has distinct requirements. For example, each user in the list requires two additional network requests to fetch all the data displayed in the list.

For these reasons, the SwiftUI view is a generic structure that contains all the common functionality, leaving the differences to be implemented by other types represented by [Swift generics](https://matteomanferdini.com/swift-generics/) with type constraints.

```swift
struct RemoteListView<
    Model: DecodableResource & Identifiable & Hashable & Placeholded,
    Content: View,
    ViewModel: RemoteListViewModel
>: View {
    let url: DecodableURL<[Model]>
    @State private var viewModel: ViewModel?
    @Environment(NetworkController.self) private var networkController

    typealias ToggleAction = () -> Void

    var body: some View {
        RemoteList(title: url.title, items: items, isLoadingNextPage: isLoadingNextPage) { item in
            // ...
        } loadNextPage: {
            guard !(viewModel?.items.isEmpty ?? true) else { return }
            Task { try? await viewModel?.fetchNextPage() }
        }
        .loading(viewModel?.items.isEmpty ?? true, redacted: true)
        .task {
            guard viewModel == nil else { return }
            viewModel = ViewModel(url: url.url, networkLayer: networkController)
            try? await viewModel?.fetch()
        }
        .refreshable {
            try? await viewModel?.refresh()
        }
    }

    var items: [Model] {
        let placeholders: [Model] = Array(repeating: .placeholder, count: 3)
        guard let items = viewModel?.items else { return placeholders }
        return items.isEmpty ? placeholders : items
    }

    var isLoadingNextPage: Bool {
        guard let viewModel else { return false }
        return viewModel.isLoading && !viewModel.items.isEmpty
    }
}
```

This is a pretty complicated piece of code. However, for this example, I want you to focus on the view model.

As I demonstrated at the beginning of the article, much of the view functionality is delegated to a view model object that can be tested separately.

The crucial part in this example, however, is the `ViewModel` generic, which allows the parent of the `RemoteListView` to specify a `RemoteListViewModel` subclass that implements the required API calls.

```swift
@MainActor @Observable class UsersListViewModel: RemoteListViewModel {
    // ...
}

struct UsersListView: View {
    let url: DecodableURL<[User]>

    var body: some View {
        RemoteListView<
            User,
            UserRow,
            UsersListViewModel
        >(url: url) { user, toggleAction in
            UserRow(user: user, toggleFollowing: toggleAction)
        }
    }
}

@MainActor @Observable class ReposListView: RemoteListViewModel {
    // ...
}

struct ReposListView: View {
    let url: DecodableURL<[Repository]>

    var body: some View {
        RemoteListView<
            Repository,
            RepositoryRow,
            ReposListViewModel
        >(url: url) { repo, toggleAction in
            RepositoryRow(repository: repo, toggleStarring: toggleAction)
        }
    }
}
```

Whenever the app needs to display a new list of entities, most of the functionality can be reused by implementing a new `RemoteListViewModel` subclass containing the specific implementation.

## Conclusions

The MVVM pattern provides the flexibility of adding view models **when it makes sense** for maintainability, testability, and code reuse.

The arguments against using view models are weak, and **the alternatives provided often violate several software development principles**. While that is not an issue in simple apps, it presents several problems at scale.

Apple frameworks are flexible and **do not impose or tacitly endorse a particular architecture**, even though they are implicitly built over MVC, like every piece of software since the inception of the pattern.

Even with particular frameworks like SwiftData, **view models can be used in conjunction with** `@Query` stored properties in views.

Finally, strict design patterns like the MV pattern, which forcefully compress rather than expand the architectural layers of an app, strip away **the unique characteristics of view models** and create structural problems that cannot be solved by shared objects alone.

If you want to learn more about using view models and how to architect your SwiftUI apps, download my free guide below.

## SwiftUI App Architecture: Design Patterns and Best Practices

![](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

It's easy to make an app by throwing some code together. But without best practices and robust architecture, you soon end up with unmanageable spaghetti code. In this guide I'll show you how to properly structure SwiftUI apps.

Matteo has been developing apps for iOS since 2008. He has been [teaching iOS development best practices](https://matteomanferdini.com/products/) to hundreds of students since 2015 and he is the developer of [Vulcan](https://purecreek.com/vulcan), a macOS app to generate SwiftUI code. Before that he was a freelance iOS developer for small and big clients, including TomTom, Squla, Siilo, and Layar. Matteo got a master’s degree in computer science and computational logic at the University of Turin. In his spare time he dances and teaches tango.