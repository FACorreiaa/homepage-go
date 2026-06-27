---
title: "3 Key Strategies to Make SwiftUI Views More Reusable"
url: https://matteomanferdini.com/swiftui-reusable-views/
author: matteomanferdini
date: 2026-05-05
ingested_at: 2026-05-13T17:45:42Z
tags:
  - swift
  - swiftui
topic: Swift
status: uncompiled
---

Monolithic views are the most common issue in SwiftUI codebases, especially among inexperienced developers.

While the declarative nature of SwiftUI makes it possible to quickly put together complex user interfaces, it also makes it easy to create massive view bodies that soon become unmaintainable.

In this article, we will see how to modularize your code to make SwiftUI views more reusable.

![Image 1](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

## FREE GUIDE - SwiftUI App Architecture: Design Patterns and Best Practices

MVC, MVVM, and MV in SwiftUI, plus practical lessons on code encapsulation and data flow.

[DOWNLOAD THE FREE GUIDE](https://matteomanferdini.com/swiftui-reusable-views/)

## Table of contents

1.   [Monolithic SwiftUI views are hard to maintain and don’t scale](https://matteomanferdini.com/swiftui-reusable-views/#monolithic)
2.   [Strategy 1. Extracting small reusable components from a view’s long body](https://matteomanferdini.com/swiftui-reusable-views/#extraction)
3.   [Computed properties are a shortcut that creates further coupling](https://matteomanferdini.com/swiftui-reusable-views/#computed-properties)
4.   [Strategy 2. Reusing styling code across views through view modifiers](https://matteomanferdini.com/swiftui-reusable-views/#modifiers)
5.   [Strategy 3. Creating reusable containers with generics](https://matteomanferdini.com/swiftui-reusable-views/#generics)

## Monolithic SwiftUI views are hard to maintain and don’t scale

In the early stages of a project, it’s tempting to keep all UI logic in a single content view that represents an entire app screen.

However, as your app grows, these views become “monsters” that are difficult to navigate and maintain, violating the [Separation of Concerns](https://en.wikipedia.org/wiki/Separation_of_concerns) principle.

Eventually, as code accumulates in a view, you start to recognize patterns that are useful in multiple places.

This leads to copy-pasting the repeated code across different views, inevitably violating the [Don’t Repeat Yourself (DRY)](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) principle and creating a **maintenance nightmare**.

There are several drawbacks to monolithic SwiftUI views:

*   **UI inconsistency**: A change to a brand color requires manually hunting down every instance of that element.
*   **Poor readability**: Deeply nested stacks and complex logic in copy-and-paste code obscure the view’s actual intent.
*   **Difficult debugging**: Fixing a UI bug in one place doesn’t propagate fixes to other copies.
*   **Fragile scaling**: Large files increase cognitive load, making the codebase hard to understand and intimidating for new team members.

Without **reusable components**, you’re building [technical debt](https://en.wikipedia.org/wiki/Technical_debt) into your codebase that can, eventually, collapse your entire project.

That’s why breaking views into smaller, focused pieces is essential for a healthy, **modular architecture**.

The simplest way to improve code quality is **view extraction**. This involves taking a chunk of code from a massive body and moving it into its own dedicated struct.

By doing this, you turn a complex layout into a readable list of **descriptive components**. Instead of seeing 50 lines of `VStack` and `HStack`, you see a single view initialization that conveys more meaning.

### The steps to extract a reusable component from an existing view

1.   **Identify**: Find a self-contained UI element, like a list row, a card, or a custom button.
2.   **Create**: Define a new struct that conforms to the `View`protocol.
3.   **Parameterize**: Use properties with simple types to pass in only the data the subview needs to render.
4.   **Replace**: Call your new view inside the parent’s body.

In a typical _monolithic view_, styling and layout are tangled together.The layout and styling logic of a small section are mixed into the main view’s layout, causing its body to grow rapidly.

```
struct ProfileView: View {
	var body: some View {
		VStack {
			VStack(alignment: .leading) {
				Text("Jane Doe")
					.font(.headline)
				Text("iOS Developer")
					.font(.subheadline)
					.foregroundColor(.secondary)
			}
			.padding()
			.background(Color.gray.opacity(0.1))
			.cornerRadius(10)
			// Other views
		}
	}
}
```

Xcode even provides an _Extract Selection to File_ option in the _Refactor_ contextual menu to simplify the process.

[![Image 2](blob:http://localhost/f04f733369778123f1fa30c7cd510219)](https://matteomanferdini.com/wp-content/uploads/2026/05/refactor-extract-selection-to-file-menu-in-xcode.jpeg)

By extracting the UI to a **reusable view**, the parent view becomes clean and declarative. It describes **what** is on the screen, while the subview encapsulates **how it appears**.

```
struct Header: View {
	let name: String
	let role: String

	var body: some View {
		VStack(alignment: .leading) {
			Text(name)
				.font(.headline)
			Text(role)
				.font(.subheadline)
				.foregroundColor(.secondary)
		}
		.padding()
		.background(Color.gray.opacity(0.1))
		.cornerRadius(10)
	}
}

struct ProfileView: View {
	var body: some View {
		VStack {
			Header(name: "Jane Doe", role: "iOS Developer")
			// Other views
		}
	}
}
```

This is the first step toward a **modular architecture**, enforces _Separation of Concerns,_ and significantly reduces **cognitive load** for readers of your code.

Your code is now more readable and visually organized, and your views are easier to preview in Xcode. You can now focus on one small piece of the UI at a time without getting lost in the noise of the entire screen.

![Image 3](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

## FREE GUIDE - SwiftUI App Architecture: Design Patterns and Best Practices

MVC, MVVM, and MV in SwiftUI, plus practical lessons on code encapsulation and data flow.

[DOWNLOAD THE FREE GUIDE](https://matteomanferdini.com/swiftui-reusable-views/)

## Computed properties are a shortcut that creates further coupling

Using computed properties to break up a large body is a common shortcut that feels helpful but is ultimately a [false economy](https://softwareengineering.stackexchange.com/questions/19573/what-are-the-worst-false-economies-in-software-development). While it makes the main body look cleaner, it fails to address the underlying architectural problems.

When you keep UI logic inside a property, you are essentially hiding complexity rather than **decoupling** it. This approach keeps your code [highly coupled](https://en.wikipedia.org/wiki/Coupling_(computer_programming)), making it hard to maintain as the app scales.

Computed properties are not wrong in principle and have their use. However, placing code that should be extracted into a reusable view in a computed property violates these design principles:

*   **[Single-Responsibility Principle](https://en.wikipedia.org/wiki/Single-responsibility_principle)**: Your main view remains responsible for both the high-level layout and the granular details of its subcomponents.
*   **[Principle of Least Knowledge](https://en.wikipedia.org/wiki/Law_of_Demeter)**: Properties often rely on the parent’s internal state, meaning the component knows too much about its surroundings.
*   **Separation of Concerns**: The sub-component cannot exist independently; it remains trapped within the parent’s scope.
*   **Don’t Repeat Yourself**: You cannot share a computed property across different screens. To use it elsewhere, you’d have to copy and paste, creating redundant code.

Properties also hinder **SwiftUI’s diffing engine**. Because a property doesn’t have its own identity, SwiftUI may struggle to optimize re-renders as effectively as it does with a dedicated struct.

Furthermore, you lose the ability to use SwiftUI previews for that specific piece of UI. A true reusable view should be a **standalone unit** that is easy to test, preview, and move.

## Strategy 2. Reusing styling code across views through view modifiers

Creating a full custom view is not the only way to reuse code in SwiftUI. If your goal is to apply a consistent **style** or **behavior**across different view types, view modifiers are the ideal solution.

A view modifier is a Swift struct conforming to the `ViewModifier` protocol that encapsulates a group of standard modifiers into a single, reusable unit, e.g., fonts, colors, or shadows.

Reusing styling code through view modifiers has several benefits:

*   **Consistency**: You define a view style once to ensure elements across your app look exactly the same.
*   **[Encapsulation](https://en.wikipedia.org/wiki/Encapsulation_(computer_programming))**: It keeps your view’s body clean by hiding a long list of styling attributes behind a single call.
*   **Versatility**: Unlike a custom struct, a modifier can be applied to any view, whether it’s a `Text`, an `Image`, or a `Button`.
*   **Extension support**: Using `View` extensions, you can apply custom logic using clean dot syntax, making the code feel like a native part of SwiftUI.

By moving styling logic into a modifier, you can again follow the _Single Responsibility Principle_. The main view focuses on content, while the modifier focuses on appearance.

Moreover, when you have multiple elements sharing the same look, repeating modifiers is again a direct violation of the _DRY principle_.

Take, for example, this series of view modifiers to style all the titles across your app.

```
Text("Welcome Back")
	.font(.largeTitle)
	.fontWeight(.bold)
	.padding()
	.background(Color.blue)
	.foregroundColor(.white)
	.cornerRadius(12)
```

We can encapsulate the styling into a struct conforming to `ViewModifier` and create an extension on `View` for seamless use at the call site.

```
struct TitleStyle: ViewModifier {
	func body(content: Content) -> some View {
		content
			.font(.largeTitle)
			.fontWeight(.bold)
			.padding()
			.background(Color.blue)
			.foregroundColor(.white)
			.cornerRadius(12)
	}
}

extension View {
	func titleStyle() -> some View {
		modifier(TitleStyle())
	}
}

// Usage
Text("Welcome Back")
	.titleStyle()
```

### The structural edge of ViewModifier types over simple methods

While a simple method in a `View` extension would work for basic styling reuse, the `ViewModifier`protocol offers deeper integration with the SwiftUI lifecycle.

A view modifier structure has several advantages over a simple method in a `View` extension:

*   **Internal state**: A `ViewModifier` can hold property wrappers such as `@State`, `@GestureState`, or `@Animation`. This allows you to build complex behaviors without cluttering the parent view’s logic.
*   **Encapsulation**: It hides implementation details. An extension method often forces you to pass in dependencies as parameters, whereas a view modifier struct can manage its own dependencies and access the environment via the `@Environment` wrapper.
*   **Improved diffing**: SwiftUI treats a `ViewModifier` as a distinct node in the view hierarchy. This helps the rendering engine optimize updates and animations more effectively than a flat chain of individual methods.
*   **Protocol extensions and generics**: Since it is a formal type, your modifier can conform to other protocols or be passed as a generic parameter to other components.

![Image 4](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

## FREE GUIDE - SwiftUI App Architecture: Design Patterns and Best Practices

MVC, MVVM, and MV in SwiftUI, plus practical lessons on code encapsulation and data flow.

[DOWNLOAD THE FREE GUIDE](https://matteomanferdini.com/swiftui-reusable-views/)

## Strategy 3. Creating reusable containers with generics

While simple parameters work for basic views, [Swift generics](https://matteomanferdini.com/swift-generics/) offer the ultimate level of flexibility. A **generic view** serves as a structural shell that can host various content types without knowing their specific identities in advance.

This is particularly useful for building **container components** like custom cards, lists, or modal wrappers.

Generics are an advanced feature that many developers struggle to understand. However, this also means they are more powerful than simpler abstraction primitives:

*   **Universal reusability**: One component can serve multiple use cases across the entire app.
*   **Code consolidation**: It drastically reduces [boilerplate](https://en.wikipedia.org/wiki/Boilerplate_code) by eliminating near-identical type definitions.
*   **Structural integrity**: It enforces a consistent layout while allowing the inner content to remain dynamic.

Using Swift generics with the `@ViewBuilder` attribute lets you create agnostic containers. This pattern separates the container’s layout from the specific content it displays.

For example, let’s assume we have a `TextCard` that displays only text. If we want to reuse its code with an image, we must build an entirely new view that shares most of the code, violating the DRY principle.

```
struct TextCard: View {
	let title: String
	let description: String

	var body: some View {
		VStack(alignment: .leading) {
			Text(title).font(.headline)
			Text(description)
		}
		.padding()
		.background(Color.secondary.opacity(0.1))
		.cornerRadius(10)
	}
}
```

However, by creating a generic container view and using the `@ViewBuilder` attribute, the card becomes a flexible shell for any view hierarchy.

```
struct Card: View {
	let title: String
	let content: Content

	init(title: String, @ViewBuilder content: () -> Content) {
		self.title = title
		self.content = content()
	}

	var body: some View {
		VStack(alignment: .leading) {
			Text(title).font(.headline)
			content
		}
		.padding()
		.background(Color.secondary.opacity(0.1))
		.cornerRadius(10)
	}
}

// Usage
Card(title: "Settings") {
	HStack {
		Image(systemName: "gear")
		Text("Manage your account")
	}
}
```

Generic container views offer features you do not get from simpler reusable views:

*   **Composition**: You can nest any view hierarchy inside the container without modifying its source code.
*   **Flexibility**: The same container works for text, images, or complex interactive views.
*   **Decoupling**: The container doesn’t need to know **what** it is displaying, only **how** to frame it.

## Conclusions

Creating **reusable views** is the hallmark of a senior developer. It transforms a rigid, tangled UI into a living **design system** that is easy to scale and maintain.

By moving away from monolithic bodies, you ensure your codebase remains healthy and adaptable. You aren’t just writing code; you are building a **modular toolkit**.

*   **Start small**: Use view extraction to break down massive views and clarify intent.
*   **Enforce style**: Leverage view modifiers to keep your visual identity consistent and centralized.
*   **Embrace flexibility**: Apply Swift generics and the @ViewBuilder attribute to create powerful, agnostic containers.

Adopting these patterns satisfies several software design principles and keeps your view logic loosely coupled. This results in fewer bugs, faster onboarding for new team members, and a much more enjoyable development experience.

Ultimately, to create a modular architecture, [apply design patterns such as MVVM to your SwiftUI codebase](https://matteomanferdini.com/swiftui-mvvm/). One key technique is to [separate SwiftUI views into a root and a content layer](https://matteomanferdini.com/mvvm-root-view).

![Image 5](https://matteomanferdini.com/wp-content/uploads/2024/09/SwiftUI-App-Architecture-Cover.png)

## FREE GUIDE - SwiftUI App Architecture: Design Patterns and Best Practices

MVC, MVVM, and MV in SwiftUI, plus practical lessons on code encapsulation and data flow.

[DOWNLOAD THE FREE GUIDE](https://matteomanferdini.com/swiftui-reusable-views/)

![Image 6: Matteo Manferdini](blob:http://localhost/346dadc7249b4f18bcd4f1d0d397bce6)

Matteo has been developing apps for iOS since 2008. He has been [teaching iOS development best practices](https://matteomanferdini.com/products/) to hundreds of students since 2015 and he is the developer of [Vulcan](https://purecreek.com/vulcan), a macOS app to generate SwiftUI code. Before that he was a freelance iOS developer for small and big clients, including TomTom, Squla, Siilo, and Layar. Matteo got a master’s degree in computer science and computational logic at the University of Turin. In his spare time he dances and teaches tango.
