import React, { useCallback } from 'react'
import Particles from 'react-tsparticles'

// this is the new common package
import type { Engine } from 'tsparticles-engine'

// this function loads all the features contained in v1 package
import { loadFull } from 'tsparticles'

const ParticlesPage = ({ children, background }: any): JSX.Element => {
    // 確保載入所有需要的 features
    const particlesInit = useCallback(async (engine: Engine) => {
        await loadFull(engine)
    }, [])

    // 重頭戲在 options
    return (
        <div style={{ opacity: 0.7 }} className="particleWrapper">
            <Particles
                init={particlesInit}
                options={{
                    fullScreen: {
                        enable: true,
                        zIndex: 1
                    },
                    particles: {
                        number: {
                            value: 100,
                            density: {
                                enable: false,
                                value_area: 800
                            }
                        },
                        color: {
                            value: '#fff'
                        },
                        shape: {
                            type: 'polygon'
                        },
                        opacity: {
                            value: 0.8,
                            random: false,
                            anim: {
                                enable: false,
                                speed: 1,
                                opacity_min: 0.1,
                                sync: false
                            }
                        },
                        size: {
                            value: 4,
                            random: false,
                            anim: {
                                enable: false,
                                speed: 50,
                                size_min: 0.1,
                                sync: false
                            }
                        },
                        rotate: {
                            value: 0,
                            random: true,
                            direction: 'clockwise',
                            animation: {
                                enable: true,
                                speed: 5,
                                sync: false
                            }
                        },
                        line_linked: {
                            enable: true,
                            distance: 150,
                            color: '#ffffff',
                            opacity: 0.4,
                            width: 1
                        },
                        move: {
                            enable: true,
                            speed: 2,
                            direction: 'none',
                            random: false,
                            straight: false,
                            out_mode: 'out',
                            attract: {
                                enable: false,
                                rotateX: 600,
                                rotateY: 1200
                            }
                        }
                    },
                    interactivity: {
                        events: {
                            onhover: {
                                enable: true,
                                mode: ['repulse']
                            },
                            onclick: {
                                enable: false,
                                mode: 'push'
                            },
                            resize: true
                        },
                        modes: {
                            grab: {
                                distance: 400,
                                line_linked: {
                                    opacity: 1
                                }
                            },
                            bubble: {
                                distance: 400,
                                size: 90,
                                duration: 2,
                                opacity: 8,
                                speed: 3
                            },
                            repulse: {
                                distance: 200
                            },
                            push: {
                                particles_nb: 4
                            },
                            remove: {
                                particles_nb: 2
                            }
                        }
                    },
                    retina_detect: true,
                    background: {
                        ...(background || {}),
                        color: 'white',
                        image: '',
                        position: '50% 50%',
                        repeat: 'no-repeat',
                        size: 'cover'
                    }
                }}
            />
            {children}
        </div>
    )
}

export default ParticlesPage
