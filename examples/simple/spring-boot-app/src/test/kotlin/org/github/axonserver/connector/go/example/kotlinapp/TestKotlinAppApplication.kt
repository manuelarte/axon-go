package org.github.manuelarte.axongo.example

import org.springframework.boot.fromApplication


fun main(args: Array<String>) {
    fromApplication<KotlinAppApplication>().with(TestcontainersConfiguration::class).run(*args)
}
