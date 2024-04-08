#!/usr/bin/env node
const packageName = "onion-kv"
function getBinaryPath() {
    // Windows binaries end with .exe so we need to special case them.
    const binaryName = process.platform === "win32" ? `${packageName}.exe` : packageName;
    // Determine package name for this platform
    const platformSpecificPackageName = `${packageName}-${process.platform}-${process.arch}`

    try {
        console.log(`${platformSpecificPackageName}/bin/${binaryName}`)
        // Resolving will fail if the optionalDependency was not installed
        return require.resolve(`${platformSpecificPackageName}/bin/${binaryName}`)
    } catch (e) {
        return require("path").join(__dirname, "..", binaryName)
    }
}

require("child_process").execFileSync(getBinaryPath(), process.argv.slice(2), {
    stdio: "inherit",
})